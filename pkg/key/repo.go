package key

import "database/sql"

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

func (repo Repo) NameExists(name string) (bool, error) {
	var exists bool
	err := repo.db.QueryRow("SELECT COUNT(*) > 0 FROM api_keys WHERE name = $1;", name).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
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

func (repo Repo) IdExists(id int) (bool, error) {
	var exists bool
	err := repo.db.QueryRow("SELECT COUNT(*) > 0 FROM api_keys WHERE id = $1", id).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
