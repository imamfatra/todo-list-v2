package middleware

import (
	"time"
)

type Maker interface {
	CreateToken(usernamse string, userId int, duration time.Duration) (string, error)
	VarifyToken(token string, userId int) (*Payload, error)
}
