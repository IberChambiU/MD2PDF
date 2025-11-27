package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrTokenExpired       = errors.New("token expirado o invalido")          // token expired or invalid
	ErrInvalidSessionId   = errors.New("id de sesion inválido")              // invalid session id
	ErrInvalidCredentials = errors.New("credenciales inválidas")             // invalid credentials
	ErrInvalidEmail       = errors.New("email inválido")                     // invalid email
	ErrEmailEmpty         = errors.New("el email no puede estar vacío")      // email cannot be empty
	ErrUserEmpty          = errors.New("el usuario no puede estar vacío")    // user cannot be empty
	ErrInvalidPassword    = errors.New("contraseña inválida")                // invalid password
	ErrPasswordEmpty      = errors.New("la contraseña no puede estar vacía") // password cannot be empty
	ErrInvalidJson        = errors.New("error: 40")                          // error: 40
)

// Payload contains the payload data of the token
type Payload struct {
	Id        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issuedAt"`
	ExpiredAt time.Time `json:"expiredAt"`
}

func NewPayload(username string, userId uuid.UUID, duration time.Duration) (*Payload, error) {

	payload := &Payload{
		Id:        userId,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return ErrTokenExpired
	}
	return nil
}
