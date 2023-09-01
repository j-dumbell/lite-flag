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

func (repo Repo) Save(apiKey ApiKey) (ApiKey, error) {
	var id int
	err := repo.db.QueryRow("INSERT INTO api_keys (name, api_key, created_at) VALUES ($1, $2, $3) RETURNING id", apiKey.Name, apiKey.ApiKey, apiKey.CreatedAt).Scan(&id)
	if err != nil {
		return ApiKey{}, err
	}
	apiKey.Id = id
	return apiKey, nil
}

func (repo Repo) FindAll() ([]ApiKey, error) {
	rows, err := repo.db.Query("SELECT id, name, api_key, created_at FROM api_keys;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var apiKeys []ApiKey
	for rows.Next() {
		var apiKey ApiKey
		if err := rows.Scan(&apiKey.Id, &apiKey.Name, &apiKey.ApiKey,
			&apiKey.CreatedAt); err != nil {
			return apiKeys, err
		}
		apiKeys = append(apiKeys, apiKey)
	}

	if err = rows.Err(); err != nil {
		return apiKeys, err
	}
	return apiKeys, nil
}

func (repo Repo) DeleteById(id int) error {
	_, err := repo.db.Exec("DELETE FROM api_keys WHERE id = $1;", id)
	return err
}

type Filters struct {
	Id     int
	Name   string
	ApiKey string
}

type Condition struct {
	field string
	value any
}

func (repo Repo) FindOneById(id int) (ApiKey, error) {
	return repo.FindOne(Filters{Id: id})
}

func (repo Repo) FindOneByName(name string) (ApiKey, error) {
	return repo.FindOne(Filters{Name: name})
}

// ToDo simplify logic
func (repo Repo) FindOne(filters Filters) (ApiKey, error) {
	query := "SELECT id, name, api_key, created_at FROM api_keys"
	conditionCounter := 1
	conditions := []string{}
	conditionValues := []any{}

	if filters.Id > 0 {
		conditions = append(conditions, fmt.Sprintf("id = $%d", conditionCounter))
		conditionValues = append(conditionValues, filters.Id)
		conditionCounter++
	}

	if filters.Name != "" {
		conditions = append(conditions, fmt.Sprintf("name = $%d", conditionCounter))
		conditionValues = append(conditionValues, filters.Name)
		conditionCounter++
	}

	if filters.ApiKey != "" {
		conditions = append(conditions, fmt.Sprintf("api_key = $%d", conditionCounter))
		conditionValues = append(conditionValues, filters.ApiKey)
		conditionCounter++
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	var apiKey ApiKey
	err := repo.db.QueryRow(query, conditionValues...).Scan(&apiKey.Id, &apiKey.Name, &apiKey.ApiKey,
		&apiKey.CreatedAt)
	if err != nil {
		return ApiKey{}, err
	}

	return apiKey, nil
}
