package fflag

import (
	"context"
	"errors"

	"github.com/j-dumbell/lite-flag/pkg/pg"
)

type Service struct {
	repo Repo
}

func NewService(repo Repo) Service {
	return Service{
		repo: repo,
	}
}

var ErrAlreadyExists = errors.New("a flag with that key already exists")

func (service *Service) Create(ctx context.Context, flag Flag) (Flag, error) {
	if err := flag.Validate(); err != nil {
		return Flag{}, err
	}

	flag, err := service.repo.Create(ctx, flag)
	if errors.Is(err, pg.ErrAlreadyExists) {
		return Flag{}, ErrAlreadyExists
	} else if err != nil {
		return Flag{}, err
	}

	return flag, nil
}

func (service *Service) FindOne(ctx context.Context, key string) (Flag, error) {
	flag, err := service.repo.FindOneByKey(ctx, key)
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

func (service *Service) Delete(ctx context.Context, key string) error {
	_, err := service.repo.FindOneByKey(ctx, key)
	if err == pg.ErrNoRows {
		return ErrNotFound
	}

	if err := service.repo.Delete(ctx, key); err != nil {
		return err
	}

	return nil
}

func (service *Service) Update(ctx context.Context, flag Flag) (Flag, error) {
	if err := flag.Validate(); err != nil {
		return Flag{}, err
	}

	_, err := service.repo.FindOneByKey(ctx, flag.Key)
	if err == pg.ErrNoRows {
		return Flag{}, ErrNotFound
	} else if err != nil {
		return Flag{}, err
	}

	return service.repo.Update(ctx, flag)
}
