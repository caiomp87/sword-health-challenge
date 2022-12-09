package interfaces

import (
	"context"
	"database/sql"

	"github.com/caiomp87/sword-health-challenge/models"
)

type ITask interface {
	Create(ctx context.Context, task *models.Task) error
	FindAll(ctx context.Context) ([]*models.Task, error)
	FindAllByUserID(ctx context.Context, userID string) ([]*models.Task, error)
	FindByID(ctx context.Context, id string) (*models.Task, error)
	FindByIDAndUserID(ctx context.Context, id, userID string) (*models.Task, error)
	UpdateByID(ctx context.Context, id, userID string, task *models.Task) error
	DeleteByID(ctx context.Context, id string) error
	Done(ctx context.Context, id, userID string) error
}

type IUser interface {
	Create(ctx context.Context, user *models.User) error
	FindAll(ctx context.Context) ([]*models.User, error)
	FindByID(ctx context.Context, id string) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
}

type IDatabase interface {
	Connect() (*sql.DB, error)
	Disconnect(db *sql.DB) error
}
