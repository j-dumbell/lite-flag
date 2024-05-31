package auth

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"slices"
	"unicode"

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

func (service *Service) CreateRootKey(ctx context.Context) (ApiKey, error) {
	apiKey := ApiKey{
		Name: "root",
		Key:  newKey(),
		Role: RoleRoot,
	}
	apiKey, err := service.repo.Create(ctx, apiKey)
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

func isValidName(name string) bool {
	if name == "" {
		return false
	}

	for _, r := range name {
		if !(unicode.IsLetter(r) || unicode.IsNumber(r) || r == '-' || r == '_') {
			return false
		}
	}
	return true
}

func (createApiKeyParams *CreateApiKeyParams) Validate() error {
	validationResult := validation.Result{}
	if !isValidName(createApiKeyParams.Name) {
		validationResult.AddFieldError("name", "must be non empty, and only contain letters, numbers, hyphens and underscores")
	}
	if !slices.Contains([]Role{RoleAdmin, RoleReadonly}, createApiKeyParams.Role) {
		validationResult.AddFieldError("role", "must be one of 'admin' | 'readonly'")
	}

	return validationResult.ToError()
}

func (service *Service) CreateKey(ctx context.Context, params CreateApiKeyParams) (ApiKey, error) {
	if err := params.Validate(); err != nil {
		return ApiKey{}, err
	}

	apiKey := ApiKey{
		Name: params.Name,
		Key:  newKey(),
		Role: params.Role,
	}
	apiKey, err := service.repo.Create(ctx, apiKey)
	if err == pg.ErrAlreadyExists {
		return ApiKey{}, ErrAlreadyExists
	} else if err != nil {
		return ApiKey{}, err
	}

	return apiKey, nil
}

var ErrKeyNotFound = errors.New("api key not found")

func (service *Service) FindOneByKey(ctx context.Context, key string) (ApiKeyRedacted, error) {
	apiKey, err := service.repo.FindOneByKey(ctx, key)
	if errors.Is(err, pg.ErrNoRows) {
		return ApiKeyRedacted{}, ErrKeyNotFound
	} else if err != nil {
		return ApiKeyRedacted{}, err
	}

	return apiKey, nil
}

func (service *Service) FindOneByName(ctx context.Context, name string) (ApiKeyRedacted, error) {
	apiKey, err := service.repo.FindOneByName(ctx, name)
	if err == pg.ErrNoRows {
		return ApiKeyRedacted{}, ErrKeyNotFound
	} else if err != nil {
		return ApiKeyRedacted{}, err
	}

	return apiKey, nil
}

var ErrCannotDeleteRoot = errors.New("root key cannot be deleted")

func (service *Service) DeleteByName(ctx context.Context, name string) error {
	apiKey, err := service.repo.FindOneByName(ctx, name)
	if err == pg.ErrNoRows {
		return ErrKeyNotFound
	} else if err != nil {
		return err
	}

	if apiKey.Role == RoleRoot {
		return ErrCannotDeleteRoot
	}

	return service.repo.DeleteByName(ctx, name)
}

func (service *Service) RotateKey(ctx context.Context, name string) (ApiKey, error) {
	apiKeyRedacted, err := service.FindOneByName(ctx, name)
	if err == pg.ErrNoRows {
		return ApiKey{}, ErrKeyNotFound
	} else if err != nil {
		return ApiKey{}, err
	}

	newApiKey := ApiKey{
		Name: name,
		Key:  newKey(),
		Role: apiKeyRedacted.Role,
	}

	if err := service.repo.Update(ctx, newApiKey); err != nil {
		return ApiKey{}, err
	}

	return newApiKey, nil
}
