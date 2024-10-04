package fflag

import (
	"context"
	"errors"

	"github.com/j-dumbell/lite-flag/pkg/fp"
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

// UpsertFlagParams is the set of parameters required to create a feature flag.
type UpsertFlagParams[T any] struct {
	// Key is the flag's key.  It must be unique, non-empty, and contain only
	// numbers, letters, underscores and hyphens.  It is immutable.
	Key string `json:"key"`

	// IsPublic determines whether the flag is public or not.
	IsPublic bool `json:"isPublic"`

	// Value is the flag's initial value.
	Value T `json:"value"`
}

func (service *Service) CreateStringFlag(ctx context.Context, stringFlag UpsertFlagParams[string]) (Flag, error) {
	flag := Flag{
		Key:         stringFlag.Key,
		Type:        FlagTypeString,
		IsPublic:    stringFlag.IsPublic,
		StringValue: &stringFlag.Value,
	}

	return service.Create(ctx, flag)
}

func (service *Service) CreateBooleanFlag(ctx context.Context, booleanFlag UpsertFlagParams[bool]) (Flag, error) {
	flag := Flag{
		Key:          booleanFlag.Key,
		Type:         FlagTypeBoolean,
		IsPublic:     booleanFlag.IsPublic,
		BooleanValue: &booleanFlag.Value,
	}

	return service.Create(ctx, flag)
}

func (service *Service) CreateJSONFlag(ctx context.Context, jsonFlag UpsertFlagParams[map[string]interface{}]) (Flag, error) {
	flag := Flag{
		Key:       jsonFlag.Key,
		Type:      FlagTypeJSON,
		IsPublic:  jsonFlag.IsPublic,
		JSONValue: jsonFlag.Value,
	}

	return service.Create(ctx, flag)
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

func (service *Service) FindAll(ctx context.Context, publicOnly bool) ([]Flag, error) {
	if publicOnly {
		return service.repo.Find(ctx, Filters{IsPublic: fp.ToPtr(true)})
	}
	return service.repo.Find(ctx, Filters{})
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
