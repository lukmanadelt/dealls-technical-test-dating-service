package service

import (
	"context"
	"dealls-technical-test-dating-service/internal/domain/entity"
	"dealls-technical-test-dating-service/internal/domain/repository"
)

// UserService is the interface used for the user service.
type UserService interface {
	CreateUser(ctx context.Context, user *entity.User) (int, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
}

type userService struct {
	repo repository.UserRepository
}

// NewUserService is a function used to initialize the user service implementation.
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

// CreateUser is a method for creating a user.
func (u *userService) CreateUser(ctx context.Context, user *entity.User) (int, error) {
	return u.repo.Insert(ctx, user)
}

// GetUserByEmail is a method for getting user based on email.
func (u *userService) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	return u.repo.FindByEmail(ctx, email)
}
