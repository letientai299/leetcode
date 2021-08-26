package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthenticationManager_CountUnexpiredTokens(t *testing.T) {
	am := Constructor1797(13)

	am.Renew("ajvy", 1)
	assert.Equal(t, 0, am.CountUnexpiredTokens(3))
	assert.Equal(t, 0, am.CountUnexpiredTokens(4))

	am.Generate("fuzxq", 5)
	am.Generate("izmry", 7)
	am.Renew("puv", 12)
	am.Generate("ybiqb", 13)
	am.Generate("gm", 14)

	assert.Equal(t, 4, am.CountUnexpiredTokens(15))
	assert.Equal(t, 3, am.CountUnexpiredTokens(18))
	assert.Equal(t, 3, am.CountUnexpiredTokens(19))

	am.Renew("ybiqb", 21)
	assert.Equal(t, 2, am.CountUnexpiredTokens(23))
	assert.Equal(t, 2, am.CountUnexpiredTokens(25))
	assert.Equal(t, 2, am.CountUnexpiredTokens(26))

	am.Generate("aqdm", 28)
	assert.Equal(t, 2, am.CountUnexpiredTokens(29))
}
