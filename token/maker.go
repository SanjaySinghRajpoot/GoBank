package token

import (
	"time"
)

type Maker interface {
	// CreateToken creates a new token
	CreateToken(username string, duration time.Duration) (string, error)

	VerifyToken(token string) (*Payload, error)
}
