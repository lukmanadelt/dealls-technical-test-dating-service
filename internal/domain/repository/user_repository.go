package repository

import (
	"context"

	"dealls-technical-test-dating-service/internal/domain/entity"
)

// UserRepository is the user repository interface.
type UserRepository interface {
	Insert(ctx context.Context, user *entity.User) (int, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
}
