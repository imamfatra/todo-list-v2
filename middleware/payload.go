package middleware

import (
	"errors"
	"time"
	"todo-api/helper"

	"github.com/google/uuid"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	UserId    int       `json:"userid"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(username string, userId int, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	helper.IfError(err)

	payload := &Payload{
		ID:        tokenID,
		UserId:    userId,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}

var ErrExpiredToken = errors.New("Token Has Expired")
var ErrInvalidToken = errors.New("Token is Invalid")
var ErrNotFoundUserId = errors.New("UserId Not Found")

func (p *Payload) IsValid() error {
	if time.Now().After(p.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
