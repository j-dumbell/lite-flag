package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/j-dumbell/lite-flag/internal/api"
	"github.com/j-dumbell/lite-flag/internal/fflag"
	"github.com/j-dumbell/lite-flag/pkg/pg"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

const serverAddress = ":8080"

func main() {
	db, err := pg.Connect(mkPGOptions())
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to DB")
	}
	defer db.Close()

	flagRepo := fflag.NewRepo(db)
	flagService := fflag.NewService(flagRepo)
	mux := api.New(db, flagService)

	log.Info().Msgf("starting webserver on address %s", serverAddress)
	err = http.ListenAndServe(serverAddress, mux.NewRouter())
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start webserver")
	}
}

func getEnv(envName string) string {
	envValue, exists := os.LookupEnv(envName)
	if exists == false {
		log.Fatal().Msgf("environment variable '%s' does not exist")
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
