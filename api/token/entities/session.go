package entities

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

var (
	// Mutual exclusion for the Session
	MuSession = sync.Mutex{}

	Sessions sessions
	///Errors for the Session
	ErrInvalidSession  = errors.New("sesión inválida")
	ErrInvalidToken    = errors.New("token inválido")
	ErrInvalidUserId   = errors.New("id de usuario inválido")
	ErrSessionNotFound = errors.New("sesión no encontrada")
	ErrInvalidDuration = errors.New("duración inválida")
)

type sessions []Session

type Session struct {
	Id           uuid.UUID
	RefreshToken string
	Duration     int64
}

func NewSession(id uuid.UUID, refreshToken string, duration int64) *Session {
	return &Session{
		Id:           id,
		RefreshToken: refreshToken,
		Duration:     duration,
	}
}

func (s *Session) UpdateToken(token string) error {
	if token == "" {
		return ErrInvalidToken
	}
	s.RefreshToken = token
	return nil
}
