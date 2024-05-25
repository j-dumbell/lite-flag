package auth

import (
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

func (repo *KeyRepo) Create(params CreateKeyParams) (ApiKey, error) {
	var id int
	err := repo.db.
		QueryRow(`
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

func (repo *KeyRepo) FindAll() ([]ApiKeyRedacted, error) {
	rows, err := repo.db.Query("SELECT id, name, role FROM api_keys;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return repo.parseRows(rows)
}

func (repo *KeyRepo) DeleteByID(id int) error {
	_, err := repo.db.Exec("DELETE FROM api_keys WHERE id = $1;", id)
	return err
}

func (repo *KeyRepo) FindOneByKey(key string) (ApiKeyRedacted, error) {
	rows, err := repo.db.Query("SELECT id, name, role FROM api_keys WHERE key = crypt($1, key)", key)
	if err != nil {
		return ApiKeyRedacted{}, err
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
