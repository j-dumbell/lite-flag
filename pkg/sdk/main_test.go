package sdk

import (
	"context"
	"database/sql"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/j-dumbell/lite-flag/internal/api"
	"github.com/j-dumbell/lite-flag/internal/auth"
	"github.com/j-dumbell/lite-flag/internal/bootstrap"
	"github.com/j-dumbell/lite-flag/internal/fflag"
	testutils "github.com/j-dumbell/lite-flag/internal/test_utils"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
)

var testDB *sql.DB
var flagService fflag.Service
var authService auth.Service
var testServer *httptest.Server

func TestMain(m *testing.M) {
	ctx := context.Background()

	container, err := testutils.StartTestDB(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start PG container")
	}
	defer func() {
		if err := container.Terminate(ctx); err != nil {
			log.Fatal().Err(err).Msg("failed to terminate PG container")
		}
	}()

	db, err := testutils.ConnectTestDB(ctx, container)
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
	testApi := api.New(testDB, flagService, authService)

	testServer = httptest.NewServer(testApi.NewRouter())
	defer testServer.Close()

	code := m.Run()
	os.Exit(code)
}

func resetDB(t *testing.T) {
	err := bootstrap.Truncate(context.Background(), testDB)
	require.NoError(t, err, "failed to reset DB")
}

func createAdminKey(t *testing.T) auth.ApiKey {
	apiKey, err := authService.CreateKey(context.Background(), auth.CreateApiKeyParams{
		Name: "admin-test",
		Role: auth.RoleAdmin,
	})
	require.NoError(t, err, "failed to create admin key")
	return apiKey
}
