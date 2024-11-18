package models

import (
	"github.com/phn00dev/go-task-manager/internal/constants"
	"time"
)

type User struct {
	ID           int              `json:"id"`
	Firstname    string           `json:"firstname"`
	Lastname     string           `json:"lastname"`
	Email        string           `json:"email"`
	PasswordHash string           `json:"password_hash"`
	UserStatus   constants.Status `json:"user_status"`
	CreatedAt    time.Time        `json:"created_at"`
	UpdatedAt    time.Time        `json:"updated_at"`
	LastLogin    time.Time        `json:"last_login"`
}
