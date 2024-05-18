package bootstrap

import "database/sql"

var ddls = []string{
	"DROP TABLE IF EXISTS flags CASCADE;",
	`CREATE TABLE flags (
		id          SERIAL PRIMARY KEY,
		name        VARCHAR NOT NULL UNIQUE,
		enabled     BOOLEAN NOT NULL
	);`,
}

var truncateStatements = []string{
	"TRUNCATE TABLE flags;",
}

func executeMany(db *sql.DB, statements []string) error {
	for _, statement := range statements {
		_, err := db.Exec(statement)
		if err != nil {
			return err
		}
	}
	return nil
}

func Recreate(db *sql.DB) error {
	return executeMany(db, ddls)
}

func Truncate(db *sql.DB) error {
	return executeMany(db, truncateStatements)
}
