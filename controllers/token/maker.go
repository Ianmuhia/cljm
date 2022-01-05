package token

import (
	"time"
)

var (
	TokenService Maker = &JWTMaker{}
)

// Maker is an interface for managing tokens
type Maker interface {
	// CreateToken creates a new token for a specific username and duration
	CreateToken(email string, duration time.Duration) (string, error)

	// CreateRefreshToken creates new refresh token for a specific email and duration
	CreateRefreshToken(email string, duration time.Duration) (string, error)

	// VerifyToken checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
