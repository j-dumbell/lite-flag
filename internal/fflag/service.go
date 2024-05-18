package fflag

import (
	"encoding/json"
	"strings"

	"github.com/j-dumbell/lite-flag/pkg/pg"
)

type ValidationError struct {
	FieldErrors map[string]string `json:"errors"`
}

func (validationError ValidationError) Error() string {
	jsonError, err := json.Marshal(validationError.FieldErrors)
	if err != nil {
		return ""
	}

	return string(jsonError)
}

func (validationError *ValidationError) AddFieldError(field string, error string) {
	validationError.FieldErrors[field] = error
}

func NewValidationError() ValidationError {
	return ValidationError{
		FieldErrors: map[string]string{},
	}
}

type UpsertFlagParams struct {
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
}

func (body *UpsertFlagParams) Validate() error {
	if strings.ContainsRune(body.Name, ' ') || strings.ContainsRune(body.Name, '/') {
		validationError := NewValidationError()
		validationError.AddFieldError("name", "name can only contain letters, numbers, hyphens or underscores")
		return &validationError
	}

	return nil
}

type Service struct {
	repo Repo
}

func NewService(repo Repo) Service {
	return Service{
		repo: repo,
	}
}

func (service *Service) Create(params UpsertFlagParams) (Flag, error) {
	if err := params.Validate(); err != nil {
		return Flag{}, err
	}

	return service.repo.Create(params)
}

func (service *Service) FindOne(id uint32) (Flag, error) {
	return service.repo.FindOne(id)
}

func (service *Service) FindAll() ([]Flag, error) {
	return service.repo.FindAll()
}

func (service *Service) Delete(id uint32) error {
	return service.repo.Delete(id)
}

func (service *Service) Update(flag Flag) (Flag, error) {
	upsertParams := UpsertFlagParams{
		Name:    flag.Name,
		Enabled: flag.Enabled,
	}
	if err := upsertParams.Validate(); err != nil {
		return Flag{}, err
	}

	_, err := service.repo.FindOne(flag.ID)
	if err != nil {
		return Flag{}, pg.ParseError(err)
	}

	return service.repo.Update(flag)
}
