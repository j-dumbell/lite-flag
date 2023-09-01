package main

import (
	"database/sql"
	"fmt"
	"github.com/j-dumbell/lite-flag/pkg/health"
	"github.com/j-dumbell/lite-flag/pkg/key"
	"github.com/j-dumbell/lite-flag/pkg/logger"
	_ "github.com/lib/pq"
	"net/http"
	"os"
	"strconv"
)

func main() {
	logger.Logger.Info("connecting to db")
	db, err := connectDb()
	if err != nil {
		logger.Logger.Error("failed to connect to DB", "error", err)
		panic(err)
	}
	defer db.Close()

	mux := http.NewServeMux()
	mux.Handle("/health", health.NewHandler(db))
	mux.Handle("/api-keys", key.NewHandler(db))

	logger.Logger.Info("starting websever")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		logger.Logger.Error("failed to start webserver", "error", err)
		panic(err)
	}
}

func getEnvOrPanic(envName string) string {
	envValue, exists := os.LookupEnv(envName)
	if exists == false {
		panic(fmt.Errorf("environment variable '%s' does not exist", envName))
	}
	return envValue
}

func connectDb() (*sql.DB, error) {
	host := getEnvOrPanic("DB_HOST")
	portEnv := getEnvOrPanic("DB_PORT")
	port, err := strconv.Atoi(portEnv)
	if err != nil {
		return nil, err
	}
	user := getEnvOrPanic("DB_USER")
	password := getEnvOrPanic("DB_PASSWORD")
	dbName := getEnvOrPanic("DB_NAME")

	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := sql.Open("postgres", psqlConn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
