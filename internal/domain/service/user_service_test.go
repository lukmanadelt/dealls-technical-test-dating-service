package service

import (
	"context"
	"reflect"
	"testing"
	"time"

	"dealls-technical-test-dating-service/internal/domain/entity"
	"dealls-technical-test-dating-service/internal/domain/repository"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	mockFailedUserRepository = &fakeUserRepository{
		id:  0,
		err: gorm.ErrDuplicatedKey,
	}
	mockSuccessUser = &entity.User{
		ID:        1,
		Email:     "user@email.com",
		Password:  "password",
		Name:      "User",
		BirthDate: time.Now().UTC(),
		Gender:    "MALE",
		Location:  "Indonesia",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
)

type fakeUserRepository struct {
	id   int
	user *entity.User
	err  error
}

func (f *fakeUserRepository) Insert(context.Context, *entity.User) (int, error) {
	return f.id, f.err
}

func (f *fakeUserRepository) FindByEmail(context.Context, string) (*entity.User, error) {
	return f.user, f.err
}

func TestUserService_CreateUser(t *testing.T) {
	type fields struct {
		repo repository.UserRepository
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		{
			name: "Failed",
			fields: fields{
				repo: mockFailedUserRepository,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Success",
			fields: fields{
				repo: &fakeUserRepository{
					id:  1,
					err: nil,
				},
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			u := &userService{
				repo: test.fields.repo,
			}
			got, err := u.CreateUser(context.Background(), &entity.User{})
			if (err != nil) != test.wantErr {
				t.Errorf("UserService.CreateUser() error = %v, wantErr %v", err, test.wantErr)

				return
			}
			if got != test.want {
				t.Errorf("UserService.CreateUser() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestUserService_GetUserByEmail(t *testing.T) {
	type fields struct {
		repo repository.UserRepository
	}
	tests := []struct {
		name    string
		fields  fields
		want    *entity.User
		wantErr bool
	}{
		{
			name: "Failed",
			fields: fields{
				repo: &fakeUserRepository{
					user: &entity.User{},
					err:  schema.ErrUnsupportedDataType,
				},
			},
			want:    &entity.User{},
			wantErr: true,
		},
		{
			name: "Success",
			fields: fields{
				repo: &fakeUserRepository{
					user: mockSuccessUser,
					err:  nil,
				},
			},
			want:    mockSuccessUser,
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			u := NewUserService(test.fields.repo)
			got, err := u.GetUserByEmail(context.Background(), "user@email.com")
			if (err != nil) != test.wantErr {
				t.Errorf("UserService.GetUserByEmail() error = %v, wantErr %v", err, test.wantErr)

				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("UserService.GetUserByEmail() = %v, want %v", got, test.want)
			}
		})
	}
}
