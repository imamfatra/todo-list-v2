package test_test

import (
	"testing"
	"todo-api/model"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	passwordString := model.RandomString(12)

	passwordHash, err := model.HashPassword(passwordString)
	assert.NoError(t, err)
	assert.NotEmpty(t, passwordHash)
	// fmt.Println(passwordHash)

	err = model.CheckPassword(passwordString, passwordHash)
	assert.NoError(t, err)

	wrongPassword := model.RandomString(12)
	err = model.CheckPassword(wrongPassword, passwordHash)
	assert.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	passwordHash2, err := model.HashPassword(passwordHash)
	assert.NoError(t, err)
	assert.NotEmpty(t, passwordHash)
	assert.NotEqual(t, passwordHash, passwordHash2)
}
