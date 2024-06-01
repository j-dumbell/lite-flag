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
