// Package controller handles HTTP requests, maps them to use cases, and returns responses.
package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"dealls-technical-test-dating-service/internal/domain/entity"
	"dealls-technical-test-dating-service/internal/domain/usecase"
	"dealls-technical-test-dating-service/pkg/constant"

	"github.com/emicklei/go-restful/v3"
	"gorm.io/gorm"
)

// UserController is a struct for handling HTTP requests and responses and mapping to use cases.
type UserController struct {
	userUsecase usecase.UserUsecase
}

// NewUserController is a function used to initialize the user controller.
func NewUserController(uu usecase.UserUsecase) *UserController {
	return &UserController{
		userUsecase: uu,
	}
}

// Signup is a method for signing up a user.
func (u *UserController) Signup(req *restful.Request, resp *restful.Response) {
	signupReq := &entity.UserSignupRequest{}
	err := req.ReadEntity(signupReq)
	if err != nil {
		resp.WriteError(http.StatusBadRequest, err)

		return
	}

	err = u.userUsecase.Signup(req.Request.Context(), signupReq)
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		resp.WriteError(http.StatusConflict, errors.New("Email already exists"))

		return
	}
	if err != nil {
		if strings.Contains(err.Error(), constant.InvalidRequestBody) {
			resp.WriteError(http.StatusBadRequest, err)

			return
		}

		resp.WriteError(http.StatusInternalServerError, fmt.Errorf("Failed to sign up user: %s", err.Error()))

		return
	}

	resp.WriteHeaderAndEntity(http.StatusCreated, "Sign up successful")
}

// Login is a method for logging in a user.
func (u *UserController) Login(req *restful.Request, resp *restful.Response) {
	loginReq := &entity.UserLoginRequest{}
	err := req.ReadEntity(loginReq)
	if err != nil {
		resp.WriteError(http.StatusBadRequest, err)

		return
	}

	loginResp, err := u.userUsecase.Login(req.Request.Context(), loginReq)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp.WriteError(http.StatusNotFound, errors.New("User not found"))

		return
	}
	if err != nil {
		if strings.Contains(err.Error(), constant.InvalidRequestBody) || err.Error() == constant.InvalidEmailPassword {
			resp.WriteError(http.StatusBadRequest, err)

			return
		}

		resp.WriteError(http.StatusInternalServerError, fmt.Errorf("Failed to logging in user: %s", err.Error()))

		return
	}

	resp.WriteHeaderAndEntity(http.StatusOK, loginResp)
}
