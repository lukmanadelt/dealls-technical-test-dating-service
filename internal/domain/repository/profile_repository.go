// Package repository defines repository interfaces.
package repository

import (
	"context"

	"dealls-technical-test-dating-service/internal/domain/entity"
)

// ProfileRepository is the profile repository interface.
type ProfileRepository interface {
	Insert(ctx context.Context, profile *entity.Profile) error
}
