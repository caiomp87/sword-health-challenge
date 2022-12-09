package models

import "time"

type User struct {
	ID           string    `json:"id,omitempty"`
	Name         string    `json:"name,omitempty"`
	Type         string    `json:"type,omitempty"`
	Email        string    `json:"email,omitempty"`
	Password     string    `json:"password,omitempty"`
	PasswordHash string    `json:"passwordHash,omitempty"`
	CreatedAt    time.Time `json:"createdAt,omitempty"`
}
