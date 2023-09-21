package token

import "time"

// Maker is an iterface for management of tokens
type Maker interface {
	// CreateToken creates a new token for provided username and duration
	CreateToken(username string, duration time.Duration) (string, *Payload, error)

	// VerifyToken checks the validity of the token
	VerifyToken(token string) (*Payload, error)
}
