package fflag

import (
	"database/sql"
	"fmt"
	"strings"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) Repo {
	return Repo{
		db: db,
	}
}

func (repo Repo) Save(flag Flag) (Flag, error) {
	var id int
	err := repo.db.QueryRow("INSERT INTO flags (name, created_at) VALUES ($1, $2) RETURNING id", flag.Name, flag.CreatedAt).Scan(&id)
	if err != nil {
		return Flag{}, err
	}
	flag.ID = id
	return flag, nil
}

type Filters struct {
	ID   *string
	Name *string
}

func (repo Repo) FindOne(filters Filters) (Flag, error) {
	query := "SELECT id, name, created_at FROM flags"
	conditionCounter := 1
	conditions := []string{}
	conditionValues := []any{}

	if filters.ID != nil {
		conditions = append(conditions, fmt.Sprintf("id = $%d", conditionCounter))
		conditionValues = append(conditionValues, *filters.ID)
		conditionCounter++
	}

	if filters.Name != nil {
		conditions = append(conditions, fmt.Sprintf("name = $%d", conditionCounter))
		conditionValues = append(conditionValues, *filters.Name)
		conditionCounter++
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	var flag Flag
	err := repo.db.QueryRow(query, conditionValues...).Scan(&flag.ID, &flag.Name, &flag.CreatedAt)
	if err != nil {
		return Flag{}, err
	}

	return flag, nil
}

func (repo Repo) FindOneByName(name string) (Flag, error) {
	return repo.FindOne(Filters{Name: &name})
}
