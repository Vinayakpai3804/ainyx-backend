package repository

import (
	"context"
	"time"
	db "ainyx-backend/db/sqlc"
	"ainyx-backend/internal/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, req models.CreateUserRequest) (db.User, error)
	GetUserByID(ctx context.Context, id int32) (db.User, error)
	UpdateUser(ctx context.Context, id int32, req models.UpdateUserRequest) (db.User, error)
	DeleteUser(ctx context.Context, id int32) error
	ListUsers(ctx context.Context, limit, offset int32) ([]db.User, error)
}

type userRepository struct {
	queries *db.Queries
}

func NewUserRepository(queries *db.Queries) UserRepository {
	return &userRepository{queries: queries}
}
func (r *userRepository) CreateUser(ctx context.Context, req models.CreateUserRequest) (db.User, error) {
	dob, err := time.Parse("2006-01-02", req.Dob)
	if err != nil {
		return db.User{}, err
	}
	return r.queries.CreateUser(ctx, db.CreateUserParams{
		Name: req.Name,
		Dob:  dob,
	})
}

func (r *userRepository) GetUserByID(ctx context.Context, id int32) (db.User, error) {
	return r.queries.GetUserByID(ctx, id)
}

func (r *userRepository) UpdateUser(ctx context.Context, id int32, req models.UpdateUserRequest) (db.User, error) {
	dob, err := time.Parse("2006-01-02", req.Dob)
	if err != nil {
		return db.User{}, err
	}
	return r.queries.UpdateUser(ctx, db.UpdateUserParams{
		ID:   id,
		Name: req.Name,
		Dob:  dob,
	})
}

func (r *userRepository) DeleteUser(ctx context.Context, id int32) error {
	return r.queries.DeleteUser(ctx, id)
}

func (r *userRepository) ListUsers(ctx context.Context, limit, offset int32) ([]db.User, error) {
	return r.queries.ListUsers(ctx, db.ListUsersParams{
		Limit:  limit,
		Offset: offset,
	})
}