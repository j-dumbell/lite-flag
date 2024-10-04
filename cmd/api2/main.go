package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/j-dumbell/lite-flag/internal/api2"
	"github.com/j-dumbell/lite-flag/internal/auth"
	"github.com/j-dumbell/lite-flag/internal/fflag"
	"github.com/j-dumbell/lite-flag/pkg/pg"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

const serverAddress = ":8080"

var shutdownTimeout = 10 * time.Second

func main() {
	if err := run(); err != nil {
		log.Fatal().Err(err).Send()
	}
}

//nolint:funlen
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
	handler := api2.New(db, flagService, authService)

	server := http.Server{
		Addr:              serverAddress,
		Handler:           handler,
		ReadHeaderTimeout: time.Second,
	}

	serverErrors := make(chan error)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Info().Msgf("starting webserver on address %s", serverAddress)
		err := server.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			serverErrors <- err
		}
	}()

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)
	case sig := <-sigChan:
		log.Info().Msgf("signal %s received; shutting down", sig.String())
		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		err := server.Shutdown(shutdownCtx)
		if err != nil {
			return fmt.Errorf("failed to shutdown gracefully: %w", err)
		}
		log.Info().Msg("successfully shutdown gracefully")
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
