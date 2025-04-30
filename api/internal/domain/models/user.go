package models

import "time"

type User struct {
	ID           string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewUser(email, passwordHash string) *User {
	now := time.Now()
	return &User{
		Email:        email,
		PasswordHash: passwordHash,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}
