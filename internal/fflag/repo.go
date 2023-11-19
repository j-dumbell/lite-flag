package fflag

import (
	"database/sql"

	"github.com/j-dumbell/lite-flag/pkg/pg"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) Repo {
	return Repo{
		db: db,
	}
}

func (repo *Repo) Save(flag Flag) error {
	_, err := repo.db.Exec("INSERT INTO flags (id, enabled, created_at) VALUES ($1, $2, $3)", flag.ID, flag.Enabled, flag.CreatedAt)
	return pg.ParseError(err)
}

func parseQueryResult(rows *sql.Rows) ([]Flag, error) {
	defer rows.Close()

	var flags []Flag
	for rows.Next() {
		var flag Flag

		if err := rows.Scan(&flag.ID, &flag.Enabled, &flag.CreatedAt); err != nil {
			return nil, err
		}

		flags = append(flags, flag)
	}

	return flags, nil
}

func (repo *Repo) FindAll() ([]Flag, error) {
	rows, err := repo.db.Query("SELECT id, enabled, created_at FROM flags;")
	if err != nil {
		return nil, pg.ParseError(err)
	}

	return parseQueryResult(rows)
}

func (repo *Repo) FindOne(id string) (Flag, error) {
	rows, err := repo.db.Query("SELECT id, enabled, created_at FROM flags WHERE id = $1;", id)
	if err != nil {
		return Flag{}, pg.ParseError(err)
	}

	flags, err := parseQueryResult(rows)
	if err != nil {
		return Flag{}, err
	}
	if len(flags) == 0 {
		return Flag{}, sql.ErrNoRows
	}
	return flags[0], nil
}

func (repo *Repo) Delete(id string) error {
	_, err := repo.db.Query("DELETE FROM flags WHERE id = $1;", id)
	return pg.ParseError(err)
}
