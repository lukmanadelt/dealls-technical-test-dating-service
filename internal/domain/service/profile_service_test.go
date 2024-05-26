package service

import (
	"context"
	"testing"

	"dealls-technical-test-dating-service/internal/domain/entity"
	"dealls-technical-test-dating-service/internal/domain/repository"

	"gorm.io/gorm"
)

type fakeProfileRepository struct {
	err error
}

func (f *fakeProfileRepository) Insert(context.Context, *entity.Profile) error {
	return f.err
}

func TestProfileService_CreateProfile(t *testing.T) {
	type fields struct {
		repo repository.ProfileRepository
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Failed",
			fields: fields{
				repo: &fakeProfileRepository{
					err: gorm.ErrDuplicatedKey,
				},
			},
			wantErr: true,
		},
		{
			name: "Success",
			fields: fields{
				repo: &fakeProfileRepository{
					err: nil,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewProfileService(tt.fields.repo)
			if err := p.CreateProfile(context.Background(), &entity.Profile{}); (err != nil) != tt.wantErr {
				t.Errorf("ProfileService.CreateProfile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
