package key

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncrypt(t *testing.T) {
	text := "helloWorld"
	key := []byte("example key 1234")

	encrypted, err := Encrypt(key, text)
	assert.NoError(t, err)
	assert.NotEqual(t, text, encrypted)
}

func TestDecrypt(t *testing.T) {
	key := []byte("example key 1234")
	text := "helloWorld"

	encrypted, err := Encrypt(key, text)
	require.NoError(t, err)

	decrypted, err := Decrypt(key, encrypted)
	assert.NoError(t, err)
	assert.Equal(t, text, decrypted)
}
