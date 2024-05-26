// Package service implements domain services containing business logic.
package service

import (
	"context"

	"dealls-technical-test-dating-service/internal/domain/entity"
	"dealls-technical-test-dating-service/internal/domain/repository"
)

// ProfileService is the interface used for the profile service.
type ProfileService interface {
	CreateProfile(ctx context.Context, profile *entity.Profile) error
}

type profileService struct {
	repo repository.ProfileRepository
}

// NewProfileService is a function used to initialize the profile service implementation.
func NewProfileService(repo repository.ProfileRepository) ProfileService {
	return &profileService{
		repo: repo,
	}
}

// CreateProfile is a method for creating a profile.
func (p *profileService) CreateProfile(ctx context.Context, profile *entity.Profile) error {
	return p.repo.Insert(ctx, profile)
}
