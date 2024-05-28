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

type CreateKeyParams struct {
	Name string
	Key  string
	Role Role
}

func (repo *KeyRepo) Create(ctx context.Context, params CreateKeyParams) (ApiKey, error) {
	var id int
	err := repo.db.
		QueryRowContext(ctx, `
			INSERT INTO api_keys (name, key, role) 
				VALUES ($1, crypt($2, gen_salt('bf')), $3) 
			RETURNING id`,
			params.Name, params.Key, params.Role).
		Scan(&id)
	if err != nil {
		return ApiKey{}, pg.ParseError(err)
	}

	return ApiKey{ID: id, Name: params.Name, Key: params.Key, Role: params.Role}, nil
}

func (repo *KeyRepo) Update(ctx context.Context, apiKey ApiKey) error {
	_, err := repo.db.ExecContext(ctx, `
			UPDATE api_keys
			SET 
			    name = $2,
			    key = crypt($3, gen_salt('bf')),
			    role = $4
			WHERE id = $1;`,
		apiKey.ID, apiKey.Name, apiKey.Key, apiKey.Role)
	if err != nil {
		return pg.ParseError(err)
	}

	return nil
}

func (repo *KeyRepo) parseRows(rows *sql.Rows) ([]ApiKeyRedacted, error) {
	var apiKeys []ApiKeyRedacted
	for rows.Next() {
		var apiKey ApiKeyRedacted
		if err := rows.Scan(&apiKey.ID, &apiKey.Name, &apiKey.Role); err != nil {
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
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name, role FROM api_keys;")
	if err != nil {
		return nil, pg.ParseError(err)
	}
	defer rows.Close()

	return repo.parseRows(rows)
}

func (repo *KeyRepo) DeleteByID(ctx context.Context, id int) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM api_keys WHERE id = $1;", id)
	return pg.ParseError(err)
}

func (repo *KeyRepo) FindOneByKey(ctx context.Context, key string) (ApiKeyRedacted, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name, role FROM api_keys WHERE key = crypt($1, key)", key)
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

func (repo *KeyRepo) FindOneByID(ctx context.Context, id int) (ApiKeyRedacted, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name, role FROM api_keys WHERE id = $1", id)
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
