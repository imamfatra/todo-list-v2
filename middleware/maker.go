package middleware

import (
	"time"
)

type Maker interface {
	CreateToken(usernamse string, duration time.Duration) (string, error)
	VarifyToken(token string) (*Payload, error)
}
