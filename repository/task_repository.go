package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/caiomp87/sword-health-challenge/db/sqlc"
	"github.com/caiomp87/sword-health-challenge/interfaces"
	"github.com/caiomp87/sword-health-challenge/models"
)

var TaskRepository interfaces.ITask

type taskDatabaseHelper struct {
	dbConn *sql.DB
	*sqlc.Queries
}

func NewTaskRepository(dbConn *sql.DB) interfaces.ITask {
	return &taskDatabaseHelper{
		dbConn:  dbConn,
		Queries: sqlc.New(dbConn),
	}
}

func (r *taskDatabaseHelper) Create(ctx context.Context, task *models.Task) error {
	return r.Queries.CreateTask(ctx, sqlc.CreateTaskParams{
		ID:          task.ID,
		Name:        task.Name,
		Summary:     task.Summary,
		Performed:   false,
		Createdat:   time.Now(),
		Performedat: time.Now(),
	})
}

func (r *taskDatabaseHelper) FindAll(ctx context.Context) ([]*models.Task, error) {
	tasks, err := r.Queries.FindAllTasks(ctx)
	if err != nil {
		return nil, err
	}

	output := make([]*models.Task, 0)
	for _, task := range tasks {
		output = append(output, &models.Task{
			ID:          task.ID,
			Name:        task.Name,
			Summary:     task.Summary,
			Performed:   task.Performed,
			CreatedAt:   task.Createdat,
			PerformedAt: task.Performedat,
		})
	}

	return output, nil
}

func (r *taskDatabaseHelper) FindByID(ctx context.Context, id string) (*models.Task, error) {
	task, err := r.Queries.FindTaskById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &models.Task{
		ID:          task.ID,
		Name:        task.Name,
		Summary:     task.Summary,
		Performed:   task.Performed,
		CreatedAt:   task.Createdat,
		PerformedAt: task.Performedat,
	}, nil
}

func (r *taskDatabaseHelper) UpdateByID(ctx context.Context, id string, task *models.Task) error {
	return r.Queries.UpdateTask(ctx, sqlc.UpdateTaskParams{
		ID:      id,
		Name:    task.Name,
		Summary: task.Summary,
	})
}

func (r *taskDatabaseHelper) DeleteByID(ctx context.Context, id string) error {
	return r.Queries.DeleteTask(ctx, id)
}

func (r *taskDatabaseHelper) Done(ctx context.Context, id string) error {
	return r.Queries.DoneTask(ctx, sqlc.DoneTaskParams{
		ID:          id,
		Performed:   true,
		Performedat: time.Now(),
	})
}
