package models

import "time"

type Task struct {
	ID          string
	Name        string
	Summary     string
	CreatedAt   time.Time
	PerformedAt time.Time
}

func NewTask(id, name, summary string) *Task {
	return &Task{
		ID:      id,
		Name:    name,
		Summary: summary,
	}
}
