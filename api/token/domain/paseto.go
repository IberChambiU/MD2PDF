package domain

import (
	"fmt"
	"time"

	"golang.org/x/crypto/chacha20poly1305"

	a "github.com/IberChambiU/MD2PDF/api/token/application"
	e "github.com/IberChambiU/MD2PDF/api/token/entities"
	"github.com/google/uuid"
	"github.com/o1egl/paseto"
)

// PasetoMaker is a PASETO token maker
type PasetoMaker struct {
	paseto      *paseto.V2
	symetricKey []byte
}

// NewPasetoMaker creates a new PasetoMaker
func NewPasetoMaker(symetricKey string) (a.Maker, error) {
	if len(symetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("symetric key must be %d bytes long", chacha20poly1305.KeySize)
	}
	maker := &PasetoMaker{
		paseto:      paseto.NewV2(),
		symetricKey: []byte(symetricKey),
	}
	return maker, nil
}

// CreateToken creates a new token for a specific role and duration
func (maker *PasetoMaker) CreateToken(role string, userId uuid.UUID, duration time.Duration) (string, error) {
	payload, err := e.NewPayload(role, userId, duration)
	if err != nil {
		return "", err
	}
	return maker.paseto.Encrypt(maker.symetricKey, payload, nil)
}

// VerifyToken Checks if a token is valid or not
func (maker *PasetoMaker) ValidateToken(token string) (*e.Payload, error) {
	payload := &e.Payload{}
	if err := maker.paseto.Decrypt(token, maker.symetricKey, payload, nil); err != nil {
		return nil, err
	}
	if err := payload.Valid(); err != nil {
		return nil, err
	}
	return payload, nil
}
