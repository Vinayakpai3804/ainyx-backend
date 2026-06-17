package service

import (
	"context"
	"time"
	"ainyx-backend/internal/models"
	"ainyx-backend/internal/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, req models.CreateUserRequest) (models.UserResponse, error)
	GetUserByID(ctx context.Context, id int32) (models.UserResponse, error)
	UpdateUser(ctx context.Context, id int32, req models.UpdateUserRequest) (models.UserResponse, error)
	DeleteUser(ctx context.Context, id int32) error
	ListUsers(ctx context.Context, page, limit int32) ([]models.UserResponse, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func calculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()
	if now.YearDay() < dob.YearDay() {
		age--
	}
	return age
}

func (s *userService) CreateUser(ctx context.Context, req models.CreateUserRequest) (models.UserResponse, error) {
	user, err := s.repo.CreateUser(ctx, req)
	if err != nil {
		return models.UserResponse{}, err
	}
	return models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob.Format("2006-01-02"),
	}, nil
}

func (s *userService) GetUserByID(ctx context.Context, id int32) (models.UserResponse, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return models.UserResponse{}, err
	}
	age := calculateAge(user.Dob)
	return models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob.Format("2006-01-02"),
		Age:  &age,
	}, nil
}

func (s *userService) UpdateUser(ctx context.Context, id int32, req models.UpdateUserRequest) (models.UserResponse, error) {
	user, err := s.repo.UpdateUser(ctx, id, req)
	if err != nil {
		return models.UserResponse{}, err
	}
	return models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob.Format("2006-01-02"),
	}, nil
}

func (s *userService) DeleteUser(ctx context.Context, id int32) error {
	return s.repo.DeleteUser(ctx, id)
}

func (s *userService) ListUsers(ctx context.Context, page, limit int32) ([]models.UserResponse, error) {
	offset := (page - 1) * limit
	users, err := s.repo.ListUsers(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	var responses []models.UserResponse
	for _, user := range users {
		age := calculateAge(user.Dob)
		responses = append(responses, models.UserResponse{
			ID:   user.ID,
			Name: user.Name,
			Dob:  user.Dob.Format("2006-01-02"),
			Age:  &age,
		})
	}
	return responses, nil
}