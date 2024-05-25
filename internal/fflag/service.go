package fflag

import (
	"strings"

	"github.com/j-dumbell/lite-flag/pkg/pg"
	"github.com/j-dumbell/lite-flag/pkg/validation"
)

type UpsertFlagParams struct {
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
}

func (body *UpsertFlagParams) Validate() error {
	if strings.ContainsRune(body.Name, ' ') || strings.ContainsRune(body.Name, '/') {
		validationError := validation.NewValidationError()
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
