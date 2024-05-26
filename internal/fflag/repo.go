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

func (repo *Repo) Create(params UpsertFlagParams) (Flag, error) {
	var id int
	err := repo.db.QueryRow("INSERT INTO flags (name, enabled) VALUES ($1, $2) RETURNING id", params.Name, params.Enabled).Scan(&id)
	if err != nil {
		return Flag{}, pg.ParseError(err)
	}

	flag := Flag{
		ID:      id,
		Name:    params.Name,
		Enabled: params.Enabled,
	}
	return flag, nil
}

func (repo *Repo) Update(flag Flag) (Flag, error) {
	_, err := repo.db.Exec("UPDATE flags SET name = $1, enabled = $2 WHERE id = $3", flag.Name, flag.Enabled, flag.ID)
	if err != nil {
		return Flag{}, pg.ParseError(err)
	}

	updatedFlag := Flag{
		ID:      flag.ID,
		Name:    flag.Name,
		Enabled: flag.Enabled,
	}

	return updatedFlag, nil
}

func parseQueryResult(rows *sql.Rows) ([]Flag, error) {
	defer rows.Close()

	var flags []Flag
	for rows.Next() {
		var flag Flag

		if err := rows.Scan(&flag.ID, &flag.Name, &flag.Enabled); err != nil {
			return nil, err
		}

		flags = append(flags, flag)
	}

	return flags, nil
}

func (repo *Repo) FindAll() ([]Flag, error) {
	rows, err := repo.db.Query("SELECT id, name, enabled FROM flags;")
	if err != nil {
		return nil, pg.ParseError(err)
	}

	return parseQueryResult(rows)
}

func (repo *Repo) FindOne(id int) (Flag, error) {
	rows, err := repo.db.Query("SELECT id, name, enabled FROM flags WHERE id = $1;", id)
	if err != nil {
		return Flag{}, pg.ParseError(err)
	}

	flags, err := parseQueryResult(rows)
	if err != nil {
		return Flag{}, err
	}
	if len(flags) == 0 {
		return Flag{}, pg.ErrNoRows
	}
	return flags[0], nil
}

func (repo *Repo) Delete(id int) error {
	_, err := repo.FindOne(id)
	if err != nil {
		return err
	}

	_, err = repo.db.Query("DELETE FROM flags WHERE id = $1;", id)
	return pg.ParseError(err)
}
