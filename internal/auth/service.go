package auth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"slices"

	"github.com/j-dumbell/lite-flag/pkg/pg"
	"github.com/j-dumbell/lite-flag/pkg/validation"
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

func (createApiKeyParams *CreateApiKeyParams) Validate() error {
	validationResult := validation.Result{}
	if createApiKeyParams.Name == "" {
		validationResult.AddFieldError("name", validation.IsRequiredMsg)
	}
	if !slices.Contains([]Role{RoleAdmin, RoleReadonly}, createApiKeyParams.Role) {
		validationResult.AddFieldError("role", "must be one of 'admin' | 'readonly'")
	}

	return validationResult.ToError()
}

func (service *Service) CreateKey(params CreateApiKeyParams) (ApiKey, error) {
	if err := params.Validate(); err != nil {
		return ApiKey{}, err
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

func (service *Service) FindOneByKey(key string) (ApiKeyRedacted, error) {
	apiKey, err := service.repo.FindOneByKey(key)
	if errors.Is(err, pg.ErrNoRows) {
		return ApiKeyRedacted{}, ErrKeyNotFound
	} else if err != nil {
		return ApiKeyRedacted{}, err
	}

	return apiKey, nil
}

func (service *Service) FindOneByID(id int) (ApiKeyRedacted, error) {
	apiKey, err := service.repo.FindOneByID(id)
	if err == pg.ErrNoRows {
		return ApiKeyRedacted{}, ErrKeyNotFound
	} else if err != nil {
		return ApiKeyRedacted{}, err
	}

	return apiKey, nil
}

var ErrCannotDeleteRoot = errors.New("root key cannot be deleted")

func (service *Service) DeleteByID(id int) error {
	apiKey, err := service.repo.FindOneByID(id)
	if err == pg.ErrNoRows {
		return ErrKeyNotFound
	} else if err != nil {
		return err
	}

	if apiKey.Role == RoleRoot {
		return ErrCannotDeleteRoot
	}

	return service.repo.DeleteByID(id)
}

func (service *Service) RotateKey(id int) (ApiKey, error) {
	apiKeyRedacted, err := service.FindOneByID(id)
	if err == pg.ErrNoRows {
		return ApiKey{}, ErrKeyNotFound
	} else if err != nil {
		return ApiKey{}, err
	}

	newApiKey := ApiKey{
		ID:   id,
		Name: apiKeyRedacted.Name,
		Key:  newKey(),
		Role: apiKeyRedacted.Role,
	}

	if err := service.repo.Update(newApiKey); err != nil {
		return ApiKey{}, err
	}

	return newApiKey, nil
}
