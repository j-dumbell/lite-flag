package fflag

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

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
		`INSERT INTO flags (key, type, is_public, boolean_value, string_value, json_value) VALUES ($1, $2, $3, $4, $5, $6)`,
		flag.Key,
		flag.Type,
		flag.IsPublic,
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
		`UPDATE flags 
					SET type = $1, is_public = $2, boolean_value = $3, string_value = $4, json_value = $5
				WHERE key = $6`,
		flag.Type,
		flag.IsPublic,
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
		if err := rows.Scan(&flag.Key, &flag.Type, &flag.IsPublic, &booleanValue, &stringValue, &jsonValue); err != nil {
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

func (repo *Repo) FindOneByKey(ctx context.Context, key string) (Flag, error) {
	return repo.FindOne(ctx, Filters{Key: &key})
}

func (repo *Repo) Delete(ctx context.Context, key string) error {
	_, err := repo.db.QueryContext(ctx, "DELETE FROM flags WHERE key = $1;", key)
	return pg.ParseError(err)
}

type Filters struct {
	Key      *string
	IsPublic *bool
}

func (repo *Repo) Find(ctx context.Context, filters Filters) ([]Flag, error) {
	conditions := []string{}
	args := []any{}
	argIndex := 1

	if filters.Key != nil {
		conditions = append(conditions, fmt.Sprintf("key = $%d", argIndex))
		args = append(args, *filters.Key)
		argIndex++
	}
	if filters.IsPublic != nil {
		conditions = append(conditions, fmt.Sprintf("is_public = $%d", argIndex))
		args = append(args, *filters.IsPublic)
		argIndex++
	}

	query := fmt.Sprintf("SELECT key, type, is_public, boolean_value, string_value, json_value FROM flags")
	if len(conditions) > 0 {
		query += fmt.Sprintf(" WHERE %s", strings.Join(conditions, " AND "))
	}

	rows, err := repo.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, pg.ParseError(err)
	}

	return parseRows(rows)
}

func (repo *Repo) FindOne(ctx context.Context, filters Filters) (Flag, error) {
	flags, err := repo.Find(ctx, filters)
	if err != nil {
		return Flag{}, err
	}
	if len(flags) == 0 {
		return Flag{}, pg.ErrNoRows
	}
	return flags[0], nil
}
