// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package sqlc

import (
	"database/sql"
)

type Task struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Summary     string       `json:"summary"`
	Createdat   sql.NullTime `json:"createdat"`
	Performedat sql.NullTime `json:"performedat"`
}