package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/j-dumbell/lite-flag/internal/auth"
	"github.com/j-dumbell/lite-flag/internal/bootstrap"
	"github.com/j-dumbell/lite-flag/pkg/pg"
	"github.com/rs/zerolog/log"
)

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

	ctx := context.Background()

	log.Info().Msg("creating tables")
	if err := bootstrap.Recreate(ctx, db); err != nil {
		return fmt.Errorf("failed to create DB tables: %w", err)
	}

	keyRepo := auth.NewKeyRepo(db)
	authService := auth.NewService(keyRepo)

	log.Info().Msg("creating root user")
	apiKey, err := authService.CreateRootKey(ctx)
	if err != nil {
		return fmt.Errorf("failed to create root user: %w", err)
	}

	log.Info().Str("API key", apiKey.Key).Msg("successfully created root user")
	return nil
}

func getEnv(envName string) (string, error) {
	envValue, exists := os.LookupEnv(envName)
	if !exists {
		return "", fmt.Errorf("environment variable '%s' does not exist", envName)
	}
	return envValue, nil
}
