package testutils

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/j-dumbell/lite-flag/pkg/pg"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	pgImage    = "docker.io/postgres:16.3"
	dbName     = "postgres"
	dbUser     = "postgres"
	dbPassword = "postgres"
)

func StartTestDB(ctx context.Context) (*postgres.PostgresContainer, error) {
	return postgres.RunContainer(ctx,
		testcontainers.WithImage(pgImage),
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
}

func ConnectTestDB(ctx context.Context, pgContainer *postgres.PostgresContainer) (*sql.DB, error) {
	host, err := pgContainer.Host(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get PG container host: %w", err)
	}

	port, err := pgContainer.MappedPort(ctx, "5432")
	if err != nil {
		return nil, fmt.Errorf("failed to get PG container port: %w", err)
	}

	return pg.Connect(pg.ConnOptions{
		DBName:   dbName,
		Username: dbUser,
		Password: dbPassword,
		Host:     host,
		Port:     port.Int(),
		SSLMode:  false,
	})
}
