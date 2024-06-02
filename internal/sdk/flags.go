package sdk

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/j-dumbell/lite-flag/internal/api"
	"github.com/j-dumbell/lite-flag/internal/fflag"
	"github.com/j-dumbell/lite-flag/pkg/fp"
)

// GetFlags returns all feature flags.
func (client *Client) GetFlags(ctx context.Context) ([]fflag.Flag, error) {
	return get[[]fflag.Flag](ctx, *client, "flags")
}

// GetFlag returns a single feature flag by its key.
func (client *Client) GetFlag(ctx context.Context, key string) (fflag.Flag, error) {
	return get[fflag.Flag](ctx, *client, buildURL("flags", key))
}

var ErrFlagValueNil = errors.New("flag value is nil")

var ErrNotStringFlag = errors.New("flag is not a string flag")

// GetStringFlagValue returns the value associated with a string feature flag.
// If the flag is not a string flag, an error is returned.
func (client *Client) GetStringFlagValue(ctx context.Context, key string) (string, error) {
	flag, err := client.GetFlag(ctx, key)
	if err != nil {
		return "", err
	}

	if flag.Type != fflag.FlagTypeString {
		return "", ErrNotStringFlag
	}

	if flag.StringValue == nil {
		return "", ErrFlagValueNil
	}

	return *flag.StringValue, nil
}

var ErrNotBooleanFlag = errors.New("flag is not a boolean flag")

// GetBooleanFlagValue returns the value associated with a boolean feature flag.
// If the flag is not a boolean flag, an error is returned.
func (client *Client) GetBooleanFlagValue(ctx context.Context, key string) (bool, error) {
	flag, err := client.GetFlag(ctx, key)
	if err != nil {
		return false, err
	}

	if flag.Type != fflag.FlagTypeBoolean {
		return false, ErrNotBooleanFlag
	}

	if flag.BooleanValue == nil {
		return false, ErrFlagValueNil
	}

	return *flag.BooleanValue, nil
}

var ErrNotJSONFlag = errors.New("flag is not a JSON flag")

// GetJSONFlagValue fetches the value associated with a JSON feature flag and unmarshals it
// into v. If the flag is not a JSON flag, an error is returned.
func (client *Client) GetJSONFlagValue(ctx context.Context, key string, v any) error {
	flag, err := client.GetFlag(ctx, key)
	if err != nil {
		return err
	}

	if flag.Type != fflag.FlagTypeJSON {
		return ErrNotStringFlag
	}

	jsonBytes, err := json.Marshal(flag.JSONValue)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(jsonBytes, v); err != nil {
		return err
	}

	return nil
}

func (client *Client) createFlag(ctx context.Context, flag fflag.Flag) error {
	if err := flag.Validate(); err != nil {
		return err
	}

	_, err := post[fflag.Flag](ctx, *client, "flags", flag)
	return err
}

// ToDo - add type-safe methods for flag-types?
func (client *Client) UpdateFlag(ctx context.Context, flag fflag.Flag) error {
	if err := flag.Validate(); err != nil {
		return err
	}

	body := api.PutFlagBody{
		Type:         flag.Type,
		IsPublic:     flag.IsPublic,
		BooleanValue: flag.BooleanValue,
		StringValue:  flag.StringValue,
		JSONValue:    flag.JSONValue,
	}

	_, err := put[fflag.Flag](ctx, *client, buildURL("flags", flag.Key), body)
	return err
}

// UpsertFlagParams is the set of parameters required to create a feature flag.
type UpsertFlagParams[T any] struct {
	// Key is the flag's key.  It must be unique, non-empty, and contain only
	// numbers, letters, underscores and hyphens.  It is immutable.
	Key string

	// IsPublic determines whether the flag is public or not.
	IsPublic bool

	// Value is the flag's value.
	Value T
}

type StringFlag = UpsertFlagParams[string]
type BooleanFlag = UpsertFlagParams[bool]
type JSONFlag = UpsertFlagParams[any]

// CreateStringFlag creates a new string feature flag.
func (client *Client) CreateStringFlag(ctx context.Context, params StringFlag) error {
	flag := fflag.Flag{
		Key:         params.Key,
		Type:        fflag.FlagTypeString,
		IsPublic:    params.IsPublic,
		StringValue: fp.ToPtr(params.Value),
	}
	return client.createFlag(ctx, flag)
}

// CreateBooleanFlag creates a new boolean feature flag.
func (client *Client) CreateBooleanFlag(ctx context.Context, params BooleanFlag) error {
	flag := fflag.Flag{
		Key:          params.Key,
		Type:         fflag.FlagTypeBoolean,
		IsPublic:     params.IsPublic,
		BooleanValue: fp.ToPtr(params.Value),
	}
	return client.createFlag(ctx, flag)
}

// CreateJSONFlag creates a new JSON feature flag.
func (client *Client) CreateJSONFlag(ctx context.Context, params JSONFlag) error {
	jsonValue, err := toJsonMap(params.Value)
	if err != nil {
		return fmt.Errorf("failed to serialize flag value as JSON: %w", err)
	}

	flag := fflag.Flag{
		Key:       params.Key,
		Type:      fflag.FlagTypeJSON,
		IsPublic:  params.IsPublic,
		JSONValue: jsonValue,
	}
	return client.createFlag(ctx, flag)
}
