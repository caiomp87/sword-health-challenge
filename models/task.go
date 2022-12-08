package models

import "time"

type Task struct {
	ID          string
	Name        string
	Summary     string
	CreatedAt   time.Time
	PerformedAt time.Time
}
