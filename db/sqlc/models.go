// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package sqlc

import (
	"time"
)

type Task struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Summary     string    `json:"summary"`
	Performed   bool      `json:"performed"`
	Createdat   time.Time `json:"createdat"`
	Performedat time.Time `json:"performedat"`
	UserID      string    `json:"user_id"`
}

type User struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Type         string    `json:"type"`
	Email        string    `json:"email"`
	Passwordhash string    `json:"passwordhash"`
	Createdat    time.Time `json:"createdat"`
}
