// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package db

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `json:"id"`
	Username       string    `json:"username"`
	HashedPassword string    `json:"hashed_password"`
	Email          string    `json:"email"`
	IsVerified     bool      `json:"is_verified"`
	IsSuperuser    bool      `json:"is_superuser"`
	Thumbnail      string    `json:"thumbnail"`
	UpdatedAt      time.Time `json:"updated_at"`
	CreatedAt      time.Time `json:"created_at"`
}

type VerifyEmail struct {
	ID         int64     `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	Email      string    `json:"email"`
	SecretCode string    `json:"secret_code"`
	IsUsed     bool      `json:"is_used"`
	CreatedAt  time.Time `json:"created_at"`
	ExpiredAt  time.Time `json:"expired_at"`
}
