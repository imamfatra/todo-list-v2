package test_test

import (
	"fmt"
	"testing"
	"time"
	"todo-api/middleware"
	"todo-api/model"

	"github.com/stretchr/testify/assert"
)

func TestTokenMaker(t *testing.T) {
	maker, err := middleware.NewPasetoMaker(model.RandomString(32))
	assert.NoError(t, err)

	username := model.RandomString(10)
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := time.Now().Add(duration)

	token, err := maker.CreateToken(username, duration)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	// fmt.Println(token)

	payload, err := maker.VarifyToken(token)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	assert.NotZero(t, payload.ID)
	assert.Equal(t, username, payload.Username)
	assert.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	assert.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)

}

func TestExpiredToken(t *testing.T) {
	maker, err := middleware.NewPasetoMaker(model.RandomString(32))
	assert.NoError(t, err)

	username := model.RandomString(10)
	duration := -time.Minute

	token, err := maker.CreateToken(username, duration)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	payload, err := maker.VarifyToken(token)
	assert.Error(t, err)
	assert.EqualError(t, err, middleware.ErrExpiredToken.Error())
	assert.Nil(t, payload)
	fmt.Println(err)
}

func TestInvalidToken(t *testing.T) {
	maker, err := middleware.NewPasetoMaker(model.RandomString(32))
	assert.NoError(t, err)

	username := model.RandomString(10)
	duration := -time.Minute

	token, err := maker.CreateToken(username, duration)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	token = token + "abc123"

	payload, err := maker.VarifyToken(token)
	assert.Error(t, err)
	assert.EqualError(t, err, middleware.ErrInvalidToken.Error())
	assert.Nil(t, payload)
	fmt.Println(err)
}
