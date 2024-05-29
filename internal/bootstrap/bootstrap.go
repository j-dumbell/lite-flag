package bootstrap

import (
	"context"
	"database/sql"
)

var ddls = []string{
	"CREATE EXTENSION IF NOT EXISTS pgcrypto;",

	"DROP TYPE IF EXISTS role CASCADE;",
	"CREATE TYPE role AS ENUM ('root', 'admin', 'readonly');",

	"DROP TYPE IF EXISTS flag_type CASCADE;",
	"CREATE TYPE flag_type AS ENUM ('string', 'boolean', 'json');",

	"DROP TABLE IF EXISTS api_keys;",
	`CREATE TABLE api_keys (
		id      SERIAL PRIMARY KEY,
		name    VARCHAR UNIQUE,
		key     VARCHAR NOT NULL UNIQUE,
		role    role NOT NULL
	);`,

	"DROP TABLE IF EXISTS flags CASCADE;",
	`CREATE TABLE flags (
		key         	VARCHAR PRIMARY KEY,
		type        	flag_type NOT NULL,
		is_public 		BOOLEAN NOT NULL,
		string_value 	VARCHAR,
		boolean_value	BOOLEAN,
		json_value		JSONB
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
