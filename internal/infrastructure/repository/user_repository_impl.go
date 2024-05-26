package repository

import (
	"context"

	"dealls-technical-test-dating-service/internal/domain/entity"

	"gorm.io/gorm"
)

// UserRepositoryImpl is a struct used to implement the user repository interface defined in the domain.
type UserRepositoryImpl struct {
	db *gorm.DB
}

// NewUserRepository is a function used to initialize the user repository implementation.
func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

// Insert is a method for inserting user data in the users table.
func (u *UserRepositoryImpl) Insert(ctx context.Context, user *entity.User) (int, error) {
	result := u.db.WithContext(ctx).Where(entity.User{Email: user.Email}).FirstOrCreate(user)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, gorm.ErrDuplicatedKey
	}

	return user.ID, nil
}

// FindByEmail is a method for finding user data based on email.
func (u *UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	user := &entity.User{}
	err := u.db.WithContext(ctx).First(user, "email = ?", email).Error

	return user, err
}
