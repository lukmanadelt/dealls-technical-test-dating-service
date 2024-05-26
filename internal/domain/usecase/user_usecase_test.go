package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"dealls-technical-test-dating-service/internal/domain"
	"dealls-technical-test-dating-service/internal/domain/entity"
	"dealls-technical-test-dating-service/internal/domain/service"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	mockPassword             = "password"
	mockSuccessSignupRequest = &entity.UserSignupRequest{
		Email:     "user@email.com",
		Password:  mockPassword,
		Name:      "name",
		BirthDate: "2024-01-01",
		Gender:    "MALE",
		Location:  "Indonesia",
	}
	mockSuccessLoginRequest = &entity.UserLoginRequest{
		Email:    "user@email.com",
		Password: mockPassword,
	}
	mockSuccessUserService = &fakeUserService{
		user: &entity.User{
			ID:        1,
			Email:     "user@email.com",
			Password:  mockPassword,
			Name:      "User",
			BirthDate: time.Now().UTC(),
			Gender:    "MALE",
			Location:  "Indonesia",
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		},
		err: nil,
	}
	mockIsValidPasswordHash = func(string, string) bool {
		return true
	}
)

type fakeUserService struct {
	id   int
	user *entity.User
	err  error
}

func (f *fakeUserService) CreateUser(context.Context, *entity.User) (int, error) {
	return f.id, f.err
}

func (f *fakeUserService) GetUserByEmail(context.Context, string) (*entity.User, error) {
	return f.user, f.err
}

type fakeProfileService struct {
	err error
}

func (f *fakeProfileService) CreateProfile(context.Context, *entity.Profile) error {
	return f.err
}

type fakeAuth struct {
	token string
	err   error
}

func (f *fakeAuth) GenerateToken(string, string) (string, error) {
	return f.token, f.err
}

type fakeConfig struct {
	key string
}

func (f *fakeConfig) GetJWTKey() string {
	return f.key
}

func Test_userUsecase_Signup(t *testing.T) {
	type fields struct {
		userService         service.UserService
		profileService      service.ProfileService
		config              domain.Config
		auth                domain.Auth
		hashPassword        hashPassword
		isValidPasswordHash isValidPasswordHash
	}
	type args struct {
		req *entity.UserSignupRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Failed: Invalid request",
			args: args{
				req: &entity.UserSignupRequest{
					Email: "",
				},
			},
			wantErr: true,
		},
		{
			name: "Failed: Password hash failed",
			fields: fields{
				hashPassword: func(string) (string, error) {
					return "", bcrypt.ErrPasswordTooLong
				},
			},
			args: args{
				req: mockSuccessSignupRequest,
			},
			wantErr: true,
		},
		{
			name: "Failed: Create user failed",
			fields: fields{
				userService: &fakeUserService{
					id:  0,
					err: gorm.ErrDuplicatedKey,
				},
				hashPassword: func(string) (string, error) {
					return mockPassword, nil
				},
			},
			args: args{
				req: mockSuccessSignupRequest,
			},
			wantErr: true,
		},
		{
			name: "Failed: Create profile failed",
			fields: fields{
				userService: &fakeUserService{
					id:  1,
					err: nil,
				},
				profileService: &fakeProfileService{
					err: gorm.ErrDuplicatedKey,
				},
				hashPassword: func(string) (string, error) {
					return mockPassword, nil
				},
			},
			args: args{
				req: mockSuccessSignupRequest,
			},
			wantErr: true,
		},
		{
			name: "Success",
			fields: fields{
				userService: &fakeUserService{
					id:  1,
					err: nil,
				},
				profileService: &fakeProfileService{
					err: nil,
				},
				hashPassword: func(string) (string, error) {
					return mockPassword, nil
				},
			},
			args: args{
				req: mockSuccessSignupRequest,
			},
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			usecase := &userUsecase{
				userService:         test.fields.userService,
				profileService:      test.fields.profileService,
				config:              test.fields.config,
				auth:                test.fields.auth,
				hashPassword:        test.fields.hashPassword,
				isValidPasswordHash: test.fields.isValidPasswordHash,
			}
			if err := usecase.Signup(context.Background(), test.args.req); (err != nil) != test.wantErr {
				t.Errorf("userUsecase.Signup() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}

func Test_userUsecase_Login(t *testing.T) {
	type fields struct {
		userService         service.UserService
		profileService      service.ProfileService
		config              domain.Config
		auth                domain.Auth
		hashPassword        hashPassword
		isValidPasswordHash isValidPasswordHash
	}
	type args struct {
		req *entity.UserLoginRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.UserLoginResponse
		wantErr bool
	}{
		{
			name: "Failed: Invalid request",
			args: args{
				req: &entity.UserLoginRequest{
					Email: "",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Failed: Get user by email failed",
			fields: fields{
				userService: &fakeUserService{
					user: &entity.User{},
					err:  schema.ErrUnsupportedDataType,
				},
			},
			args: args{
				req: mockSuccessLoginRequest,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Failed: Invalid password hash",
			fields: fields{
				userService: mockSuccessUserService,
				isValidPasswordHash: func(string, string) bool {
					return false
				},
			},
			args: args{
				req: mockSuccessLoginRequest,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Failed: Generate token failed",
			fields: fields{
				userService: mockSuccessUserService,
				auth: &fakeAuth{
					err: errors.New("Failed to generate a token"),
				},
				config:              &fakeConfig{},
				isValidPasswordHash: mockIsValidPasswordHash,
			},
			args: args{
				req: mockSuccessLoginRequest,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Success",
			fields: fields{
				userService: mockSuccessUserService,
				auth: &fakeAuth{
					token: "token",
					err:   nil,
				},
				config:              &fakeConfig{},
				isValidPasswordHash: mockIsValidPasswordHash,
			},
			args: args{
				req: mockSuccessLoginRequest,
			},
			want:    entity.NewUserLoginResponse("token"),
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			u := NewUserUsecase(test.fields.userService, test.fields.profileService, test.fields.config, test.fields.auth, test.fields.hashPassword, test.fields.isValidPasswordHash)
			got, err := u.Login(context.Background(), test.args.req)
			if (err != nil) != test.wantErr {
				t.Errorf("userUsecase.Login() error = %v, wantErr %v", err, test.wantErr)

				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("userUsecase.Login() = %v, want %v", got, test.want)
			}
		})
	}
}
