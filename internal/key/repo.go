package key

import (
	"database/sql"
	"fmt"
	"strings"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) Repo {
	return Repo{db: db}
}

func (repo *Repo) Save(apiKey ApiKey) error {
	_, err := repo.db.Exec("INSERT INTO api_keys (id, key, role, created_at) VALUES ($1, $2, $3, $4)", apiKey.ID, apiKey.ApiKey, apiKey.Role, apiKey.CreatedAt)
	return err
}

func (repo *Repo) FindAll() ([]ApiKey, error) {
	rows, err := repo.db.Query("SELECT id, key, role, created_at FROM api_keys;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var apiKeys []ApiKey
	for rows.Next() {
		var apiKey ApiKey
		if err := rows.Scan(&apiKey.ID, &apiKey.ApiKey, &apiKey.Role, &apiKey.CreatedAt); err != nil {
			return apiKeys, err
		}
		apiKeys = append(apiKeys, apiKey)
	}

	if err = rows.Err(); err != nil {
		return apiKeys, err
	}

	return apiKeys, nil
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
		conditions = append(conditions, fmt.Sprintf("api_key = $%d", conditionCounter))
		conditionValues = append(conditionValues, *filters.ApiKey)
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

	var apiKey ApiKey
	err := repo.db.QueryRow(query, conditionValues...).Scan(&apiKey.ID, &apiKey.ApiKey,
		&apiKey.Role, &apiKey.CreatedAt)
	if err != nil {
		return ApiKey{}, err
	}

	return apiKey, nil
}
