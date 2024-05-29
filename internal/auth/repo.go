package auth

import (
	"context"
	"database/sql"

	"github.com/j-dumbell/lite-flag/pkg/pg"
)

type KeyRepo struct {
	db *sql.DB
}

func NewKeyRepo(db *sql.DB) KeyRepo {
	return KeyRepo{db: db}
}

func (repo *KeyRepo) Create(ctx context.Context, apiKey ApiKey) (ApiKey, error) {
	_, err := repo.db.
		ExecContext(ctx, `
			INSERT INTO api_keys (name, key, role) 
				VALUES ($1, crypt($2, gen_salt('bf')), $3)`,
			apiKey.Name, apiKey.Key, apiKey.Role)
	if err != nil {
		return ApiKey{}, pg.ParseError(err)
	}

	return apiKey, nil
}

func (repo *KeyRepo) Update(ctx context.Context, apiKey ApiKey) error {
	_, err := repo.db.ExecContext(ctx, `
			UPDATE api_keys
			SET 
			    key = crypt($1, gen_salt('bf')),
			    role = $2
			WHERE name = $3;`,
		apiKey.Key, apiKey.Role, apiKey.Name)
	if err != nil {
		return pg.ParseError(err)
	}

	return nil
}

func (repo *KeyRepo) parseRows(rows *sql.Rows) ([]ApiKeyRedacted, error) {
	var apiKeys []ApiKeyRedacted
	for rows.Next() {
		var apiKey ApiKeyRedacted
		if err := rows.Scan(&apiKey.Name, &apiKey.Role); err != nil {
			return apiKeys, err
		}

		apiKeys = append(apiKeys, apiKey)
	}

	if err := rows.Err(); err != nil {
		return apiKeys, err
	}

	return apiKeys, nil
}

func (repo *KeyRepo) FindAll(ctx context.Context) ([]ApiKeyRedacted, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT name, role FROM api_keys;")
	if err != nil {
		return nil, pg.ParseError(err)
	}
	defer rows.Close()

	return repo.parseRows(rows)
}

func (repo *KeyRepo) DeleteByName(ctx context.Context, name string) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM api_keys WHERE name = $1;", name)
	return pg.ParseError(err)
}

func (repo *KeyRepo) FindOneByKey(ctx context.Context, key string) (ApiKeyRedacted, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT name, role FROM api_keys WHERE key = crypt($1, key)", key)
	if err != nil {
		return ApiKeyRedacted{}, pg.ParseError(err)
	}
	defer rows.Close()

	apiKeys, err := repo.parseRows(rows)
	if err != nil {
		return ApiKeyRedacted{}, err
	}

	if len(apiKeys) == 0 {
		return ApiKeyRedacted{}, pg.ErrNoRows
	}

	return apiKeys[0], nil
}

func (repo *KeyRepo) FindOneByName(ctx context.Context, name string) (ApiKeyRedacted, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT name, role FROM api_keys WHERE name = $1", name)
	if err != nil {
		return ApiKeyRedacted{}, pg.ParseError(err)
	}
	defer rows.Close()

	apiKeys, err := repo.parseRows(rows)
	if err != nil {
		return ApiKeyRedacted{}, err
	}

	if len(apiKeys) == 0 {
		return ApiKeyRedacted{}, pg.ErrNoRows
	}

	return apiKeys[0], nil
}
