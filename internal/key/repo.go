package key

import (
	"database/sql"
	"fmt"
	"strings"
)

type Repo struct {
	db        *sql.DB
	cipherKey []byte
}

func NewRepo(db *sql.DB, cipherKey []byte) Repo {
	return Repo{db: db, cipherKey: cipherKey}
}

func (repo *Repo) Save(apiKey ApiKey) error {
	encryptedKey, err := Encrypt(repo.cipherKey, apiKey.ApiKey)
	if err != nil {
		return err
	}

	_, err = repo.db.Exec("INSERT INTO api_keys (id, key, role, created_at) VALUES ($1, $2, $3, $4)", apiKey.ID, encryptedKey, apiKey.Role, apiKey.CreatedAt)
	return err
}

func (repo *Repo) parseRows(rows *sql.Rows) ([]ApiKey, error) {
	var apiKeys []ApiKey
	for rows.Next() {
		var apiKey ApiKey
		var encryptedKey string
		if err := rows.Scan(&apiKey.ID, &encryptedKey, &apiKey.Role, &apiKey.CreatedAt); err != nil {
			return apiKeys, err
		}

		decryptedApiKey, err := Decrypt(repo.cipherKey, encryptedKey)
		if err != nil {
			return apiKeys, err
		}
		apiKey.ApiKey = decryptedApiKey
		apiKeys = append(apiKeys, apiKey)
	}

	if err := rows.Err(); err != nil {
		return apiKeys, err
	}

	return apiKeys, nil
}

func (repo *Repo) FindAll() ([]ApiKey, error) {
	rows, err := repo.db.Query("SELECT id, key, role, created_at FROM api_keys;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return repo.parseRows(rows)
}

func (repo *Repo) DeleteByID(id string) error {
	_, err := repo.db.Exec("DELETE FROM api_keys WHERE id = $1;", id)
	return err
}

type Filters struct {
	ID     *string
	ApiKey *string
	Role   Role
}

func (repo *Repo) FindOneByID(id string) (ApiKey, error) {
	return repo.FindOne(Filters{ID: &id})
}

func (repo *Repo) FindOneByKey(key string) (ApiKey, error) {
	return repo.FindOne(Filters{ApiKey: &key})
}

// ToDo simplify logic
func (repo *Repo) FindOne(filters Filters) (ApiKey, error) {
	query := "SELECT id, key, role, created_at FROM api_keys"
	conditionCounter := 1
	conditions := []string{}
	conditionValues := []any{}

	if filters.ID != nil {
		conditions = append(conditions, fmt.Sprintf("id = $%d", conditionCounter))
		conditionValues = append(conditionValues, *filters.ID)
		conditionCounter++
	}

	if filters.ApiKey != nil {
		encryptedKey, err := Encrypt(repo.cipherKey, *filters.ApiKey)
		if err != nil {
			return ApiKey{}, err
		}
		conditions = append(conditions, fmt.Sprintf("api_key = $%d", conditionCounter))
		conditionValues = append(conditionValues, encryptedKey)
		conditionCounter++
	}

	if filters.Role != "" {
		conditions = append(conditions, fmt.Sprintf("role = $%d", conditionCounter))
		conditionValues = append(conditionValues, filters.Role)
		conditionCounter++
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	rows, err := repo.db.Query(query, conditionValues...)
	if err != nil {
		return ApiKey{}, err
	}
	defer rows.Close()

	apiKeys, err := repo.parseRows(rows)
	if len(apiKeys) == 0 {
		return ApiKey{}, sql.ErrNoRows
	}

	return apiKeys[0], nil
}
