package main

import (
	"os"
	"strconv"

	"github.com/j-dumbell/lite-flag/internal/auth"
	"github.com/j-dumbell/lite-flag/internal/bootstrap"
	"github.com/j-dumbell/lite-flag/pkg/pg"
	"github.com/rs/zerolog/log"
)

func main() {
	db, err := pg.Connect(mkPGOptions())
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to DB")
	}
	defer db.Close()

	log.Info().Msg("creating tables")
	if err := bootstrap.Recreate(db); err != nil {
		log.Fatal().Err(err).Msg("failed to create DB tables")
	}

	keyRepo := auth.NewKeyRepo(db)
	authService := auth.NewService(keyRepo)

	log.Info().Msg("creating root user")
	apiKey, err := authService.CreateRootKey()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create root user")
	}

	log.Info().Str("API key", apiKey.Key).Msg("successfully created root user")
}

func getEnv(envName string) string {
	envValue, exists := os.LookupEnv(envName)
	if exists == false {
		log.Fatal().Msgf("environment variable '%s' does not exist", envName)
	}
	return envValue
}

func mkPGOptions() pg.ConnOptions {
	host := getEnv("DB_HOST")
	portEnv := getEnv("DB_PORT")
	port, err := strconv.Atoi(portEnv)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to convert port to integer")
	}
	user := getEnv("DB_USER")
	password := getEnv("DB_PASSWORD")
	dbName := getEnv("DB_NAME")

	return pg.ConnOptions{
		DBName:   dbName,
		Username: user,
		Password: password,
		Host:     host,
		Port:     port,
		SSLMode:  false,
	}
}
