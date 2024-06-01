package sdk

import (
	"context"
	"encoding/json"
	"errors"

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

// CreateFlagParams is the set of parameters required to create a feature flag.
type CreateFlagParams[T any] struct {
	// Key is the flag's key.  It must be unique, non-empty, and contain only
	// numbers, letters, underscores and hyphens.
	Key string

	// IsPublic determines whether the flag is public or not.
	IsPublic bool

	// Value is the flag's initial value.
	Value T
}

// CreateStringFlag creates a new string feature flag.
func (client *Client) CreateStringFlag(ctx context.Context, params CreateFlagParams[string]) error {
	flag := fflag.Flag{
		Key:         params.Key,
		Type:        fflag.FlagTypeString,
		IsPublic:    params.IsPublic,
		StringValue: fp.ToPtr(params.Value),
	}
	return client.createFlag(ctx, flag)
}

// CreateBooleanFlag creates a new boolean feature flag.
func (client *Client) CreateBooleanFlag(ctx context.Context, params CreateFlagParams[bool]) error {
	flag := fflag.Flag{
		Key:          params.Key,
		Type:         fflag.FlagTypeBoolean,
		IsPublic:     params.IsPublic,
		BooleanValue: fp.ToPtr(params.Value),
	}
	return client.createFlag(ctx, flag)
}

// CreateJSONFlag creates a new JSON feature flag.
func (client *Client) CreateJSONFlag(ctx context.Context, params CreateFlagParams[map[string]interface{}]) error {
	flag := fflag.Flag{
		Key:       params.Key,
		Type:      fflag.FlagTypeJSON,
		IsPublic:  params.IsPublic,
		JSONValue: params.Value,
	}
	return client.createFlag(ctx, flag)
}
