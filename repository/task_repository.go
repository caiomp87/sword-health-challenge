package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/caiomp87/sword-health-challenge/interfaces"
	"github.com/caiomp87/sword-health-challenge/models"
	"github.com/caiomp87/sword-health-challenge/sql/sqlc"
)

type TaskRepository struct {
	dbConn *sql.DB
	*sqlc.Queries
}

func NewTaskRepository(dbConn *sql.DB) interfaces.ITask {
	return &TaskRepository{
		dbConn:  dbConn,
		Queries: sqlc.New(dbConn),
	}
}

func (r *TaskRepository) Create(ctx context.Context, task *models.Task) error {
	return r.Queries.CreateTask(ctx, sqlc.CreateTaskParams{
		ID:      task.ID,
		Name:    task.Name,
		Summary: task.Summary,
	})
}

func (r *TaskRepository) FindAll(ctx context.Context) ([]*models.Task, error) {
	tasks, err := r.Queries.FindAllTasks(ctx)
	if err != nil {
		return nil, err
	}

	var output []*models.Task
	for _, task := range tasks {
		output = append(output, &models.Task{
			ID:          task.ID,
			Name:        task.Name,
			Summary:     task.Summary,
			CreatedAt:   task.Createdat.Time,
			PerformedAt: task.Performedat.Time,
		})
	}

	return output, nil
}

func (r *TaskRepository) FindByID(ctx context.Context, id string) (*models.Task, error) {
	task, err := r.Queries.FindTaskById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &models.Task{
		ID:          task.ID,
		Name:        task.Name,
		Summary:     task.Summary,
		CreatedAt:   task.Createdat.Time,
		PerformedAt: task.Performedat.Time,
	}, nil
}

func (r *TaskRepository) UpdateByID(ctx context.Context, id string, task *models.Task) error {
	return r.Queries.UpdateTask(ctx, sqlc.UpdateTaskParams{
		ID:      id,
		Name:    task.Name,
		Summary: task.Summary,
	})
}

func (r *TaskRepository) DeleteByID(ctx context.Context, id string) error {
	return r.Queries.DeleteTask(ctx, id)
}

func (r *TaskRepository) Done(ctx context.Context, id string) error {
	return r.Queries.DoneTask(ctx, sqlc.DoneTaskParams{
		ID: id,
		Performedat: sql.NullTime{
			Time: time.Now(),
		},
	})
}
