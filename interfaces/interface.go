package interfaces

import (
	"context"
	"database/sql"

	"github.com/caiomp87/sword-health-challenge/models"
)

type ITask interface {
	Create(ctx context.Context, task *models.Task) error
	FindAll(ctx context.Context) ([]*models.Task, error)
	FindByID(ctx context.Context, id string) (*models.Task, error)
	UpdateByID(ctx context.Context, id string, task *models.Task) error
	DeleteByID(ctx context.Context, id string) error
	Done(ctx context.Context, id string) error
}

type IDatabase interface {
	Connect() (*sql.DB, error)
	Disconnect(db *sql.DB) error
}
