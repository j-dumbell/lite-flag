package api

import (
	"context"
	"database/sql"
	"os"
	"testing"
	"time"

	"github.com/j-dumbell/lite-flag/internal/bootstrap"
	"github.com/j-dumbell/lite-flag/internal/fflag"
	"github.com/j-dumbell/lite-flag/pkg/pg"
	"github.com/rs/zerolog/log"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

var testDB *sql.DB
var testApi API
var flagService fflag.Service

func TestMain(m *testing.M) {
	ctx := context.Background()

	dbName := "postgres"
	dbUser := "postgres"
	dbPassword := "postgres"

	postgresContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("docker.io/postgres:16.3"),
		postgres.WithDatabase("postgres"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start Postgres container")
	}

	defer func() {
		if err := postgresContainer.Terminate(ctx); err != nil {
			log.Fatal().Err(err).Msg("failed to terminate Postgres container")
		}
	}()

	host, err := postgresContainer.Host(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get Postgres container host")
	}

	port, err := postgresContainer.MappedPort(ctx, "5432")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get Postgres container port")
	}

	db, err := pg.Connect(pg.ConnOptions{
		DBName:   dbName,
		Username: dbUser,
		Password: dbPassword,
		Host:     host,
		Port:     port.Int(),
		SSLMode:  false,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to DB")
	}

	testDB = db
	if err := bootstrap.Recreate(testDB); err != nil {
		log.Fatal().Err(err).Msg("failed to create DB schema")
	}

	flagRepo := fflag.NewRepo(testDB)
	flagService = fflag.NewService(flagRepo)
	testApi = New(testDB, flagService)

	code := m.Run()
	os.Exit(code)
}
