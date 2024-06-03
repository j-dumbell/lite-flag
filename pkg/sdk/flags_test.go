package sdk

import (
	"context"
	"testing"

	"github.com/j-dumbell/lite-flag/internal/fflag"
	"github.com/j-dumbell/lite-flag/pkg/fp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_GetFlags(t *testing.T) {
	resetDB(t)
	key := createAdminKey(t)
	client := NewClient(testServer.URL, &key.Key)

	flag1, err := flagService.Create(context.Background(), fflag.Flag{
		Key:          "abc",
		Type:         fflag.FlagTypeBoolean,
		IsPublic:     true,
		BooleanValue: fp.ToPtr(true),
	})
	require.NoError(t, err, "failed to create test flag")

	flag2, err := flagService.Create(context.Background(), fflag.Flag{
		Key:         "de-f",
		Type:        fflag.FlagTypeString,
		IsPublic:    false,
		StringValue: fp.ToPtr("foobar"),
	})
	require.NoError(t, err, "failed to create test flag")

	flags, err := client.GetFlags(context.Background())
	require.NoError(t, err, "GetFlags should not error")
	assert.ElementsMatch(t, []fflag.Flag{flag1, flag2}, flags)
}

func TestClient_GetFlags_none(t *testing.T) {
	resetDB(t)
	key := createAdminKey(t)
	client := NewClient(testServer.URL, &key.Key)

	flags, err := client.GetFlags(context.Background())
	require.NoError(t, err, "GetFlags should not error")
	assert.ElementsMatch(t, []fflag.Flag{}, flags)
}

func TestClient_GetFlag(t *testing.T) {
	resetDB(t)
	key := createAdminKey(t)
	client := NewClient(testServer.URL, &key.Key)

	expectedFlag, err := flagService.Create(context.Background(), fflag.Flag{
		Key:          "abc",
		Type:         fflag.FlagTypeBoolean,
		IsPublic:     true,
		BooleanValue: fp.ToPtr(true),
	})
	require.NoError(t, err, "failed to create test flag")

	actualFlag, err := client.GetFlag(context.Background(), expectedFlag.Key)
	require.NoError(t, err, "GetFlag should not error")
	assert.Equal(t, expectedFlag, actualFlag)
}

func TestClient_GetStringFlagValue(t *testing.T) {
	resetDB(t)
	key := createAdminKey(t)
	client := NewClient(testServer.URL, &key.Key)

	flag := fflag.Flag{
		Key:         "abc",
		Type:        fflag.FlagTypeString,
		IsPublic:    false,
		StringValue: fp.ToPtr("blah"),
	}
	_, err := flagService.Create(context.Background(), flag)
	require.NoError(t, err, "failed to create test flag")

	actual, err := client.GetStringFlagValue(context.Background(), flag.Key)
	require.NoError(t, err, "GetStringFlagValue should not error")
	assert.Equal(t, *flag.StringValue, actual)
}

func TestClient_GetBooleanFlagValue(t *testing.T) {
	resetDB(t)
	key := createAdminKey(t)
	client := NewClient(testServer.URL, &key.Key)

	flag := fflag.Flag{
		Key:          "abc",
		Type:         fflag.FlagTypeBoolean,
		IsPublic:     true,
		BooleanValue: fp.ToPtr(true),
	}
	_, err := flagService.Create(context.Background(), flag)
	require.NoError(t, err, "failed to create test flag")

	actual, err := client.GetBooleanFlagValue(context.Background(), flag.Key)
	require.NoError(t, err, "GetBooleanFlagValue should not error")
	assert.Equal(t, *flag.BooleanValue, actual)
}

func TestClient_GetJSONFlagValue(t *testing.T) {
	resetDB(t)
	key := createAdminKey(t)
	client := NewClient(testServer.URL, &key.Key)

	flag := fflag.Flag{
		Key:      "abc",
		Type:     fflag.FlagTypeJSON,
		IsPublic: true,
		JSONValue: map[string]interface{}{
			"foo": 1.1,
			"bar": "yo",
			"baz": true,
		},
	}
	_, err := flagService.Create(context.Background(), flag)
	require.NoError(t, err, "failed to create test flag")

	var actual map[string]interface{}
	err = client.GetJSONFlagValue(context.Background(), flag.Key, &actual)
	require.NoError(t, err, "GetJSONFlagValue should not error")
	assert.Equal(t, flag.JSONValue, actual)
}

func TestClient_CreateBooleanFlag(t *testing.T) {
	resetDB(t)
	key := createAdminKey(t)
	client := NewClient(testServer.URL, &key.Key)

	err := client.CreateBooleanFlag(context.Background(), BooleanFlag{
		Key:      "123-4",
		IsPublic: false,
		Value:    true,
	})
	require.NoError(t, err, "CreateBooleanFlag should not error")
}

func TestClient_CreateStringFlag(t *testing.T) {
	resetDB(t)
	key := createAdminKey(t)
	client := NewClient(testServer.URL, &key.Key)

	err := client.CreateStringFlag(context.Background(), StringFlag{
		Key:      "123-4-5",
		IsPublic: false,
		Value:    "blahblah",
	})
	require.NoError(t, err, "CreateBooleanFlag should not error")
}

func TestClient_CreateJSONFlag(t *testing.T) {
	resetDB(t)
	key := createAdminKey(t)
	client := NewClient(testServer.URL, &key.Key)

	err := client.CreateJSONFlag(context.Background(), JSONFlag{
		Key:      "123-4-5",
		IsPublic: false,
		Value: map[string]interface{}{
			"foo": "bar",
		},
	})
	require.NoError(t, err, "CreateJSONFlag should not error")
}

func TestClient_UpdateFlag(t *testing.T) {
	resetDB(t)
	key := createAdminKey(t)
	client := NewClient(testServer.URL, &key.Key)

	expectedFlag, err := flagService.Create(context.Background(), fflag.Flag{
		Key:          "abc",
		Type:         fflag.FlagTypeBoolean,
		IsPublic:     true,
		BooleanValue: fp.ToPtr(true),
	})
	require.NoError(t, err, "failed to create test flag")

	err = client.UpdateFlag(context.Background(), fflag.Flag{
		Key:         expectedFlag.Key,
		Type:        fflag.FlagTypeString,
		IsPublic:    false,
		StringValue: fp.ToPtr("abc"),
	})
	require.NoError(t, err, "UpdateFlag should not error")
}
