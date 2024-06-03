package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/j-dumbell/lite-flag/internal/api"
	"github.com/j-dumbell/lite-flag/internal/auth"
	"github.com/j-dumbell/lite-flag/internal/fflag"
	"github.com/j-dumbell/lite-flag/pkg/pg"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

const serverAddress = ":8080"

func main() {
	if err := run(); err != nil {
		log.Fatal().Err(err)
	}
}

func run() error {
	host, err := getEnv("DB_HOST")
	if err != nil {
		return err
	}

	portEnv, err := getEnv("DB_PORT")
	if err != nil {
		return err
	}
	port, err := strconv.Atoi(portEnv)
	if err != nil {
		return fmt.Errorf("unable to convert port to integer: %w", err)
	}

	user, err := getEnv("DB_USER")
	if err != nil {
		return err
	}

	password, err := getEnv("DB_PASSWORD")
	if err != nil {
		return err
	}

	dbName, err := getEnv("DB_NAME")
	if err != nil {
		return err
	}

	db, err := pg.Connect(pg.ConnOptions{
		DBName:   dbName,
		Username: user,
		Password: password,
		Host:     host,
		Port:     port,
		SSLMode:  false,
	})
	if err != nil {
		return fmt.Errorf("failed to connect to DB: %w", err)
	}
	defer db.Close()

	flagRepo := fflag.NewRepo(db)
	flagService := fflag.NewService(flagRepo)
	keyRepo := auth.NewKeyRepo(db)
	authService := auth.NewService(keyRepo)
	mux := api.New(db, flagService, authService)

	log.Info().Msgf("starting webserver on address %s", serverAddress)
	err = http.ListenAndServe(serverAddress, mux.NewRouter())
	if err != nil {
		return fmt.Errorf("failed to start webserver: %w", err)
	}

	return nil
}

func getEnv(envName string) (string, error) {
	envValue, exists := os.LookupEnv(envName)
	if !exists {
		return "", fmt.Errorf("environment variable '%s' does not exist", envName)
	}
	return envValue, nil
}
