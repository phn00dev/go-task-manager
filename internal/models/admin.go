package models

import "time"

type Admin struct {
	ID           int        `json:"id"`
	Firstname    string     `json:"firstname"`
	Lastname     string     `json:"lastname"`
	Email        string     `json:"email"`
	PasswordHash string     `json:"-"`
	AdminRole    string     `json:"admin_role"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	LastLogin    *time.Time `json:"last_login"`
}
