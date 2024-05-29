package fflag

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/j-dumbell/lite-flag/pkg/fp"
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

func (repo *Repo) Create(ctx context.Context, flag Flag) (Flag, error) {
	var jsonValue *string
	if flag.JSONValue != nil {
		bytes, err := json.Marshal(flag.JSONValue)
		if err != nil {
			return Flag{}, fmt.Errorf("failed to marshal json %w", err)
		}

		jsonValue = fp.ToPtr(string(bytes))
	}

	_, err := repo.db.ExecContext(
		ctx,
		`INSERT INTO flags (key, type, boolean_value, string_value, json_value) VALUES ($1, $2, $3, $4, $5)`,
		flag.Key,
		flag.Type,
		pg.ToNullBool(flag.BooleanValue),
		pg.ToNullString(flag.StringValue),
		jsonValue,
	)
	if err != nil {
		return Flag{}, pg.ParseError(err)
	}

	return flag, nil
}

func (repo *Repo) Update(ctx context.Context, flag Flag) (Flag, error) {
	var jsonValue []byte
	if flag.JSONValue != nil {
		bytes, err := json.Marshal(flag.JSONValue)
		if err != nil {
			return Flag{}, fmt.Errorf("failed to marshal json %w", err)
		}

		jsonValue = bytes
	}

	_, err := repo.db.ExecContext(
		ctx,
		`UPDATE flags 
					SET type = $1, boolean_value = $2, string_value = $3, json_value = $4 
				WHERE key = $5`,
		flag.Type,
		pg.ToNullBool(flag.BooleanValue),
		pg.ToNullString(flag.StringValue),
		jsonValue,
		flag.Key,
	)
	if err != nil {
		return Flag{}, pg.ParseError(err)
	}

	return flag, nil
}

func parseRows(rows *sql.Rows) ([]Flag, error) {
	defer rows.Close()

	flags := []Flag{}
	for rows.Next() {
		var flag Flag

		var stringValue sql.NullString
		var booleanValue sql.NullBool
		var jsonValue sql.NullString
		if err := rows.Scan(&flag.Key, &flag.Type, &booleanValue, &stringValue, &jsonValue); err != nil {
			return nil, err
		}

		flag.StringValue = pg.FromNullString(stringValue)
		flag.BooleanValue = pg.FromNullBool(booleanValue)
		if jsonValue.Valid {
			err := json.Unmarshal([]byte(jsonValue.String), &flag.JSONValue)
			if err != nil {
				return nil, err
			}
		}

		flags = append(flags, flag)
	}

	return flags, nil
}

func (repo *Repo) FindAll(ctx context.Context) ([]Flag, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT key, type, boolean_value, string_value, json_value FROM flags;")
	if err != nil {
		return nil, pg.ParseError(err)
	}

	return parseRows(rows)
}

func (repo *Repo) FindOneByKey(ctx context.Context, key string) (Flag, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT key, type, boolean_value, string_value, json_value FROM flags WHERE key = $1;", key)
	if err != nil {
		return Flag{}, pg.ParseError(err)
	}

	flags, err := parseRows(rows)
	if err != nil {
		return Flag{}, err
	}
	if len(flags) == 0 {
		return Flag{}, pg.ErrNoRows
	}
	return flags[0], nil
}

func (repo *Repo) Delete(ctx context.Context, key string) error {
	_, err := repo.db.QueryContext(ctx, "DELETE FROM flags WHERE key = $1;", key)
	return pg.ParseError(err)
}
