package pg

import (
	"database/sql"
	"fmt"
)

// ConnOptions contains the parameters required to connect to a Postgres DB.
type ConnOptions struct {
	DBName   string
	Username string
	Password string
	Host     string
	Port     int
	SSLMode  bool
}

// Connect connects to a Postgres DB using the provided connection options.
func Connect(options ConnOptions) (*sql.DB, error) {
	var sslMode string
	if options.SSLMode {
		sslMode = "enable"
	} else {
		sslMode = "disable"
	}

	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", options.Host, options.Port, options.Username, options.Password, options.DBName, sslMode)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
