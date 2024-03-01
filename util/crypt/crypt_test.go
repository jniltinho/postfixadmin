package crypt_test

import (
	"testing"

	"postfixadmin/util/crypt"

	_ "postfixadmin/util/crypt/apr1_crypt"

	"github.com/stretchr/testify/assert"
)

func TestIsHashSupported(t *testing.T) {
	apr1 := crypt.IsHashSupported("$apr1$salt$hash")
	assert.True(t, apr1)
	other := crypt.IsHashSupported("$unknown$salt$hash")
	assert.False(t, other)
}
