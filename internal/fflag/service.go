package fflag

import (
	"errors"

	"github.com/j-dumbell/lite-flag/pkg/pg"
	"github.com/j-dumbell/lite-flag/pkg/validation"
)

type UpsertFlagParams struct {
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
}

func (upsertFlagParams *UpsertFlagParams) Validate() error {
	result := validation.Result{}
	if upsertFlagParams.Name == "" {
		result.AddFieldError("name", validation.IsRequiredMsg)
	}

	return result.ToError()
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

func (service *Service) FindOne(id int) (Flag, error) {
	return service.repo.FindOne(id)
}

func (service *Service) FindAll() ([]Flag, error) {
	return service.repo.FindAll()
}

var ErrNotFound = errors.New("no flag found with that id")

func (service *Service) Delete(id int) error {
	err := service.repo.Delete(id)
	if err == pg.ErrNoRows {
		return ErrNotFound
	} else if err != nil {
		return err
	}

	return nil
}

func (service *Service) Update(id int, upsertParams UpsertFlagParams) (Flag, error) {
	if err := upsertParams.Validate(); err != nil {
		return Flag{}, err
	}

	_, err := service.repo.FindOne(id)
	if err == pg.ErrNoRows {
		return Flag{}, ErrNotFound
	} else if err != nil {
		return Flag{}, err
	}

	flag := Flag{
		ID:      id,
		Name:    upsertParams.Name,
		Enabled: upsertParams.Enabled,
	}

	return service.repo.Update(flag)
}
