// Package usecase defines use cases (application services) that interact with the repositories and entities.
package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"dealls-technical-test-dating-service/internal/domain"
	"dealls-technical-test-dating-service/internal/domain/entity"
	"dealls-technical-test-dating-service/internal/domain/service"
	"dealls-technical-test-dating-service/pkg/constant"
)

type (
	hashPassword        func(password string) (string, error)
	isValidPasswordHash func(password, hashedPassword string) bool
)

// UserUsecase is the interface used for the user use case.
type UserUsecase interface {
	Signup(ctx context.Context, req *entity.UserSignupRequest) error
	Login(ctx context.Context, req *entity.UserLoginRequest) (*entity.UserLoginResponse, error)
}

type userUsecase struct {
	userService         service.UserService
	profileService      service.ProfileService
	config              domain.Config
	auth                domain.Auth
	hashPassword        hashPassword
	isValidPasswordHash isValidPasswordHash
}

// NewUserUsecase is a function used to initialize the user use case implementation.
func NewUserUsecase(us service.UserService, ps service.ProfileService, cfg domain.Config, a domain.Auth, hashPassword hashPassword, ivph isValidPasswordHash) UserUsecase {
	return &userUsecase{
		userService:         us,
		profileService:      ps,
		config:              cfg,
		auth:                a,
		hashPassword:        hashPassword,
		isValidPasswordHash: ivph,
	}
}

func (u *userUsecase) Signup(ctx context.Context, req *entity.UserSignupRequest) error {
	err := req.Validate()
	if err != nil {
		return fmt.Errorf("%s: %s", constant.InvalidRequestBody, err.Error())
	}

	password, err := u.hashPassword(req.Password)
	if err != nil {
		return err
	}

	currentTime := time.Now().UTC()
	user := entity.NewUser(strings.ToLower(req.Email), password, req.Name, req.BirthDate, req.Gender, req.Location, "", currentTime, currentTime)
	id, err := u.userService.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	profile := entity.NewProfile(id, "", "", false, currentTime, currentTime)
	err = u.profileService.CreateProfile(ctx, profile)
	if err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) Login(ctx context.Context, req *entity.UserLoginRequest) (*entity.UserLoginResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, fmt.Errorf("%s: %s", constant.InvalidRequestBody, err.Error())
	}

	user, err := u.userService.GetUserByEmail(ctx, strings.ToLower(req.Email))
	if err != nil {
		return nil, err
	}

	if !u.isValidPasswordHash(req.Password, user.Password) {
		return nil, errors.New(constant.InvalidEmailPassword)
	}

	token, err := u.auth.GenerateToken(user.Email, u.config.GetJWTKey())
	if err != nil {
		return nil, err
	}

	resp := entity.NewUserLoginResponse(token)

	return resp, nil
}
