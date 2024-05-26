// Package repository contains implementation of the repository interfaces defined in the domain package.
package repository

import (
	"context"

	"dealls-technical-test-dating-service/internal/domain/entity"

	"gorm.io/gorm"
)

// ProfileRepositoryImpl is a struct used to implement the profile repository interface defined in the domain.
type ProfileRepositoryImpl struct {
	db *gorm.DB
}

// NewProfileRepository is a function used to initialize the profile repository implementation.
func NewProfileRepository(db *gorm.DB) *ProfileRepositoryImpl {
	return &ProfileRepositoryImpl{
		db: db,
	}
}

// Insert is a method for inserting profile data in the profiles table.
func (p *ProfileRepositoryImpl) Insert(ctx context.Context, profile *entity.Profile) error {
	result := p.db.WithContext(ctx).Where(entity.Profile{UserID: profile.UserID}).FirstOrCreate(profile)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrDuplicatedKey
	}

	return nil
}
