package middleware

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	SymmetrycKey []byte
}

func NewPasetoMaker(symmetrycKey string) (Maker, error) {
	if len(symmetrycKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("Invalid key size. Key size must be %d", chacha20poly1305.KeySize)
	}

	key := &PasetoMaker{
		paseto:       paseto.NewV2(),
		SymmetrycKey: []byte(symmetrycKey),
	}
	return key, nil
}

func (paseto *PasetoMaker) CreateToken(usernamse string, duration time.Duration) (string, error) {
	payload, err := NewPayload(usernamse, duration)
	if err != nil {
		return "", err
	}

	return paseto.paseto.Encrypt(paseto.SymmetrycKey, payload, nil)
}

func (paseto *PasetoMaker) VarifyToken(token string) (*Payload, error) {
	var payload = &Payload{}

	err := paseto.paseto.Decrypt(token, paseto.SymmetrycKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.IsValid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
