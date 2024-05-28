package bootstrap

import (
	"context"
	"database/sql"
)

var ddls = []string{
	"CREATE EXTENSION IF NOT EXISTS pgcrypto;",

	"DROP TYPE IF EXISTS role CASCADE;",
	"CREATE TYPE role AS ENUM ('root', 'admin', 'readonly');",

	"DROP TABLE IF EXISTS api_keys;",
	`CREATE TABLE api_keys (
		id      SERIAL PRIMARY KEY,
		name    VARCHAR UNIQUE,
		key     VARCHAR NOT NULL UNIQUE,
		role    role NOT NULL
	);`,

	"DROP TABLE IF EXISTS flags CASCADE;",
	`CREATE TABLE flags (
		id          SERIAL PRIMARY KEY,
		name        VARCHAR NOT NULL UNIQUE,
		enabled     BOOLEAN NOT NULL
	);`,
}

var truncateStatements = []string{
	"TRUNCATE TABLE flags;",
	"TRUNCATE TABLE api_keys;",
}

func executeMany(ctx context.Context, db *sql.DB, statements []string) error {
	for _, statement := range statements {
		_, err := db.ExecContext(ctx, statement)
		if err != nil {
			return err
		}
	}
	return nil
}

func Recreate(ctx context.Context, db *sql.DB) error {
	return executeMany(ctx, db, ddls)
}

func Truncate(ctx context.Context, db *sql.DB) error {
	return executeMany(ctx, db, truncateStatements)
}
