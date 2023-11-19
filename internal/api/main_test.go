package api

import (
	"database/sql"
	"os"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/j-dumbell/lite-flag/internal/bootstrap"
	"github.com/j-dumbell/lite-flag/pkg/pg"
	_ "github.com/lib/pq"
)

var testDB *sql.DB
var api *chi.Mux

func TestMain(m *testing.M) {
	db, err := pg.Connect(pg.ConnOptions{
		DBName:   "postgres",
		Username: "postgres",
		Password: "postgres",
		Host:     "localhost",
		Port:     5432,
		SSLMode:  false,
	})
	if err != nil {
		panic(err)
	}
	testDB = db
	defer testDB.Close()

	err = bootstrap.Run(testDB)
	if err != nil {
		panic(err)
	}

	api = New(testDB)

	code := m.Run()
	os.Exit(code)
}
