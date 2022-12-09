package models

import "time"

type Task struct {
	ID          string    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Summary     string    `json:"summary,omitempty"`
	Performed   bool      `json:"performed"`
	UserID      string    `json:"userID"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	PerformedAt time.Time `json:"performedAt,omitempty"`
}

func NewTask(id, name, summary string) *Task {
	return &Task{
		ID:      id,
		Name:    name,
		Summary: summary,
	}
}
