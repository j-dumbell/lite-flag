package fflag

import (
	"database/sql"
	"time"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) Repo {
	return Repo{
		db: db,
	}
}

func (repo Repo) Save(flag Flag) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO flags (name, created_at) VALUES ($1, $2)", flag.Name, flag.CreatedAt)
	if err != nil {
		tx.Rollback()
		return err
	}

	if flag.Schedule == nil {
		_, err = tx.Exec("INSERT INTO transitions (flag_name, to_state, effective_from) VALUES ($1, $2, $3)", flag.Name, flag.Enabled, time.Now())
		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		_, err = tx.Exec("INSERT INTO transitions (flag_name, to_state, effective_from) VALUES ($1, $2, $3)", flag.Name, flag.Schedule.ToState, flag.Schedule.EffectiveFrom)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func parseQueryResult(rows *sql.Rows) ([]Flag, error) {
	defer rows.Close()

	var flags []Flag
	for rows.Next() {
		var flag Flag
		var scheduleToState sql.NullBool
		var scheduleEffectiveFrom sql.NullTime

		if err := rows.Scan(&flag.Name, &flag.CreatedAt, &flag.Enabled, &scheduleToState, &scheduleEffectiveFrom); err != nil {
			return nil, err
		}

		if scheduleEffectiveFrom.Valid {
			flag.Schedule = &Schedule{
				ToState:       scheduleToState.Bool,
				EffectiveFrom: scheduleEffectiveFrom.Time,
			}
		} else {
			flag.Schedule = nil
		}

		flags = append(flags, flag)
	}

	return flags, nil
}

const findQuery = `
	WITH current_state AS (
		SELECT t.flag_name, t.to_state, t.effective_from
		FROM transitions t 
		INNER JOIN (
			SELECT flag_name, MAX(effective_from) as effective_from
			FROM transitions
			WHERE effective_from <= NOW()
			GROUP BY flag_name
		) m
			ON t.flag_name = m.flag_name
			AND t.effective_from = m.effective_from
	),
	
	schedule AS (
		SELECT t.flag_name, t.to_state, t.effective_from
		FROM transitions t 
		INNER JOIN (
			SELECT flag_name, MAX(effective_from) as effective_from
			FROM transitions
			WHERE effective_from > NOW()
			GROUP BY flag_name
		) m
			ON t.flag_name = m.flag_name
			AND t.effective_from = m.effective_from
	)
	
	SELECT f.name, f.created_at, c.to_state, s.to_state, s.effective_from
	FROM flags f 
	LEFT JOIN current_state c 
		ON f.name = c.flag_name 
	LEFT JOIN schedule s
		ON f.name = s.flag_name
`

func (repo Repo) FindAll() ([]Flag, error) {
	rows, err := repo.db.Query(findQuery)
	if err != nil {
		return nil, err
	}

	return parseQueryResult(rows)
}

func (repo Repo) FindOne(name string) (Flag, error) {
	query := findQuery + " WHERE f.name = $1"
	rows, err := repo.db.Query(query, name)
	if err != nil {
		return Flag{}, err
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
