package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/caiomp87/sword-health-challenge/db/sqlc"
	"github.com/caiomp87/sword-health-challenge/interfaces"
	"github.com/caiomp87/sword-health-challenge/models"
)

var UserRepository interfaces.IUser

type userDatabaseHelper struct {
	dbConn *sql.DB
	*sqlc.Queries
}

func NewUserRepository(dbConn *sql.DB) interfaces.IUser {
	return &userDatabaseHelper{
		dbConn:  dbConn,
		Queries: sqlc.New(dbConn),
	}
}

func (r *userDatabaseHelper) Create(ctx context.Context, user *models.User) error {
	return r.Queries.CreateUser(ctx, sqlc.CreateUserParams{
		ID:           user.ID,
		Name:         user.Name,
		Type:         user.Type,
		Email:        user.Email,
		Passwordhash: user.PasswordHash,
		Createdat:    time.Now(),
	})
}

func (r *userDatabaseHelper) FindByID(ctx context.Context, id string) (*models.User, error) {
	user, err := r.Queries.FindUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:        user.ID,
		Name:      user.Name,
		Type:      user.Type,
		Email:     user.Email,
		CreatedAt: user.Createdat,
	}, nil
}

func (r *userDatabaseHelper) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := r.Queries.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:           user.ID,
		Name:         user.Name,
		Type:         user.Type,
		Email:        user.Email,
		PasswordHash: user.Passwordhash,
		CreatedAt:    user.Createdat,
	}, nil
}
