package bootstrap

import "database/sql"

var ddls = []string{
	"DROP TYPE IF EXISTS role CASCADE;",
	"CREATE TYPE role AS ENUM ('root', 'admin', 'readonly');",

	"DROP TABLE IF EXISTS api_keys CASCADE;",
	`CREATE TABLE api_keys (
		id VARCHAR PRIMARY KEY,
		key VARCHAR NOT NULL UNIQUE,
		role role NOT NULL,
		created_at TIMESTAMPTZ NOT NULL
	);`,

	"DROP TABLE IF EXISTS flags CASCADE;",
	`CREATE TABLE flags (
		id VARCHAR PRIMARY KEY,
		enabled BOOLEAN NOT NULL,
		created_at TIMESTAMPTZ NOT NULL
	);`,
}

func Run(db *sql.DB) error {
	for _, ddl := range ddls {
		_, err := db.Exec(ddl)
		if err != nil {
			return err
		}
	}
	return nil
}
