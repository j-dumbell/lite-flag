package main

import (
	"fmt"
	"github.com/j-dumbell/lite-flag/pkg/api"
	"github.com/j-dumbell/lite-flag/pkg/key"
	"github.com/j-dumbell/lite-flag/pkg/logger"
	"github.com/j-dumbell/lite-flag/pkg/pg"
	_ "github.com/lib/pq"
	"net/http"
	"os"
	"strconv"
)

func main() {
	logger.Logger.Info("connecting to db")
	db, err := pg.Connect(mkPGString())
	if err != nil {
		logger.Logger.Error("failed to connect to DB", "error", err)
		panic(err)
	}
	defer db.Close()

	rootApiKey := getEnvOrPanic("ROOT_API_KEY")
	err = key.InsertRoot(key.NewRepo(db), rootApiKey)
	if err != nil {
		logger.Logger.Error("failed to apply root API key", "error", err.Error())
		panic(err)
	}

	logger.Logger.Info("starting websever")
	mux := api.New(db)
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		logger.Logger.Error("failed to start webserver", "error", err.Error())
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

func mkPGString() pg.ConnOptions {
	host := getEnvOrPanic("DB_HOST")
	portEnv := getEnvOrPanic("DB_PORT")
	port, err := strconv.Atoi(portEnv)
	if err != nil {
		panic(err)
	}
	user := getEnvOrPanic("DB_USER")
	password := getEnvOrPanic("DB_PASSWORD")
	dbName := getEnvOrPanic("DB_NAME")

	return pg.ConnOptions{
		DBName:   dbName,
		Username: user,
		Password: password,
		Host:     host,
		Port:     port,
		SSLMode:  false,
	}
}
