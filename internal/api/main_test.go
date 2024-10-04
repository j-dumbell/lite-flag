package api

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/j-dumbell/lite-flag/internal/auth"
	"github.com/j-dumbell/lite-flag/internal/bootstrap"
	"github.com/j-dumbell/lite-flag/internal/fflag"
	"github.com/j-dumbell/lite-flag/pkg/pg"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

var testDB *sql.DB
var flagService fflag.Service
var authService auth.Service
var testHandler http.Handler

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
	if err := bootstrap.Recreate(context.Background(), testDB); err != nil {
		log.Fatal().Err(err).Msg("failed to create DB schema")
	}

	flagRepo := fflag.NewRepo(testDB)
	flagService = fflag.NewService(flagRepo)

	keyRepo := auth.NewKeyRepo(testDB)
	authService = auth.NewService(keyRepo)
	testHandler = New(testDB, flagService, authService)

	code := m.Run()
	os.Exit(code)
}

func resetDB(t *testing.T) {
	err := bootstrap.Truncate(context.Background(), testDB)
	require.NoError(t, err, "failed to reset DB")
}

func createRootKey(t *testing.T) auth.ApiKey {
	apiKey, err := authService.CreateRootKey(context.Background())
	require.NoError(t, err, "failed to create root key")
	return apiKey
}

func createAdminKey(t *testing.T) auth.ApiKey {
	apiKey, err := authService.CreateKey(context.Background(), auth.CreateApiKeyParams{
		Name: "admin-test",
		Role: auth.RoleAdmin,
	})
	require.NoError(t, err, "failed to create admin key")
	return apiKey
}

func createReadonlyKey(t *testing.T) auth.ApiKey {
	apiKey, err := authService.CreateKey(context.Background(), auth.CreateApiKeyParams{
		Name: "readonly-test",
		Role: auth.RoleReadonly,
	})
	require.NoError(t, err, "failed to create readonly key")
	return apiKey
}
