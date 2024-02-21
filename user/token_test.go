package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken_GeneratedAndValidates(t *testing.T) {
	assert := assert.New(t)
	tokenGenerator := &SimpleTokenGenerator{}

	user := &UserToken{
		Id:     1,
		Email:  "email@test.com",
		Name:   "My Name",
		Gender: "male",
		Age:    55,
	}
	content, err := tokenGenerator.Generate(user)

	assert.NoError(err)

	tokenUser, err := tokenGenerator.Validate(content)
	assert.NoError(err)
	assert.Equal(user, tokenUser)
}

func TestGenerateToken_FailsValidation(t *testing.T) {
	assert := assert.New(t)
	tokenGenerator := &SimpleTokenGenerator{}

	user, err := tokenGenerator.Validate("eyJlY2hvIjoiaGVsbG8ifQo=")

	assert.NoError(err)
	assert.Nil(user)
}

func TestGenerateToken_Validation_Errors_NonJSON(t *testing.T) {
	assert := assert.New(t)
	tokenGenerator := &SimpleTokenGenerator{}

	user, err := tokenGenerator.Validate("ZWNobwo=")

	assert.Error(err)
	assert.Nil(user)
}
