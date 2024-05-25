package auth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"

	"github.com/j-dumbell/lite-flag/pkg/pg"
)

type Service struct {
	repo KeyRepo
}

func NewService(repo KeyRepo) Service {
	return Service{
		repo: repo,
	}
}

func newKey() string {
	b := make([]byte, 40)
	rand.Read(b)
	return hex.EncodeToString(b)
}

var ErrAlreadyExists = errors.New("an API key with that name already exists")

func (service *Service) CreateRootKey() (ApiKey, error) {
	createParams := CreateKeyParams{
		Name: "root",
		Key:  newKey(),
		Role: RoleRoot,
	}
	apiKey, err := service.repo.Create(createParams)
	if err == pg.ErrAlreadyExists {
		return ApiKey{}, ErrAlreadyExists
	} else if err != nil {
		return ApiKey{}, err
	}

	return apiKey, nil
}

type CreateApiKeyParams struct {
	Name string `json:"name"`
	Role Role   `json:"role"`
}

var ErrInvalidRole = errors.New("cannot create API key with 'root' role")

func (service *Service) CreateKey(params CreateApiKeyParams) (ApiKey, error) {
	if params.Role == RoleRoot {
		return ApiKey{}, ErrInvalidRole
	}

	createParams := CreateKeyParams{
		Name: params.Name,
		Key:  newKey(),
		Role: params.Role,
	}
	apiKey, err := service.repo.Create(createParams)
	if err == pg.ErrAlreadyExists {
		return ApiKey{}, ErrAlreadyExists
	} else if err != nil {
		return ApiKey{}, err
	}

	return apiKey, nil
}

func (service *Service) KeyExists(key string) (bool, error) {
	_, err := service.repo.FindOneByKey(key)
	if err == pg.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

var ErrKeyNotFound = errors.New("api key not found")

func (service *Service) KeyRole(key string) (Role, error) {
	apiKey, err := service.repo.FindOneByKey(key)
	if errors.Is(err, pg.ErrNoRows) {
		return "", ErrKeyNotFound
	} else if err != nil {
		return "", err
	}

	return apiKey.Role, nil
}
