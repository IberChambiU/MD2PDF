package application

import (
	"time"

	e "github.com/IberChambiU/MD2PDF/api/token/entities"
	"github.com/google/uuid"
)

// Maker is an interface for managing tokens
type Maker interface {
	// CreateToken creates a new token for a specific role and duration
	CreateToken(role string, userId uuid.UUID, duration time.Duration) (string, error)
	// VerifyToken Checks if a token is valid or not
	ValidateToken(token string) (*e.Payload, error)
}
