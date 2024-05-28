package fflag

import (
	"context"
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

func (service *Service) Create(ctx context.Context, params UpsertFlagParams) (Flag, error) {
	if err := params.Validate(); err != nil {
		return Flag{}, err
	}

	flag, err := service.repo.Create(ctx, params)
	if errors.Is(err, pg.ErrAlreadyExists) {
		return Flag{}, err
	} else if err != nil {
		return Flag{}, err
	}

	return flag, nil
}

func (service *Service) FindOne(ctx context.Context, id int) (Flag, error) {
	flag, err := service.repo.FindOne(ctx, id)
	if err == pg.ErrNoRows {
		return Flag{}, ErrNotFound
	} else if err != nil {
		return Flag{}, err
	}

	return flag, nil
}

func (service *Service) FindAll(ctx context.Context) ([]Flag, error) {
	return service.repo.FindAll(ctx)
}

var ErrNotFound = errors.New("no flag found")

func (service *Service) Delete(ctx context.Context, id int) error {
	err := service.repo.Delete(ctx, id)
	if err == pg.ErrNoRows {
		return ErrNotFound
	} else if err != nil {
		return err
	}

	return nil
}

func (service *Service) Update(ctx context.Context, id int, upsertParams UpsertFlagParams) (Flag, error) {
	if err := upsertParams.Validate(); err != nil {
		return Flag{}, err
	}

	_, err := service.repo.FindOne(ctx, id)
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

	return service.repo.Update(ctx, flag)
}
