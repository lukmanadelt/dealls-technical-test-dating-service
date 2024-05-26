package integration_test

import (
	"dealls-technical-test-dating-service/internal/domain/entity"
	"dealls-technical-test-dating-service/pkg/util"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

const (
	signupURL = "/dating/v1/users/signup"
	loginURL  = "/dating/v1/users/login"
)

func TestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(Test))
}

func (t *Test) Test_Signup_Failed_Bad_Request() {
	response, err := t.executePost(signupURL, nil)
	t.Require().Equal(http.StatusBadRequest, response.Code)
	t.Require().NoError(err)
}

func (t *Test) Test_Signup_Failed_Email_Empty() {
	request := entity.UserSignupRequest{
		Email: "",
	}
	response, err := t.executePost(signupURL, request)
	t.Require().Equal(http.StatusBadRequest, response.Code)
	t.Require().NoError(err)
}

func (t *Test) Test_Signup_Failed_Invalid_Email_Format() {
	request := entity.UserSignupRequest{
		Email: "invalid@email",
	}
	response, err := t.executePost(signupURL, request)
	t.Require().Equal(http.StatusBadRequest, response.Code)
	t.Require().NoError(err)
}

func (t *Test) Test_Signup_Failed_Password_Empty() {
	request := entity.UserSignupRequest{
		Email:    "user@email.com",
		Password: "",
	}
	response, err := t.executePost(signupURL, request)
	t.Require().Equal(http.StatusBadRequest, response.Code)
	t.Require().NoError(err)
}

func (t *Test) Test_Signup_Failed_Password_Too_Short() {
	request := entity.UserSignupRequest{
		Email:    "user@email.com",
		Password: "pass",
	}
	response, err := t.executePost(signupURL, request)
	t.Require().Equal(http.StatusBadRequest, response.Code)
	t.Require().NoError(err)
}

func (t *Test) Test_Signup_Failed_Name_Empty() {
	request := entity.UserSignupRequest{
		Email:    "user@email.com",
		Password: "pass",
		Name:     "",
	}
	response, err := t.executePost(signupURL, request)
	t.Require().Equal(http.StatusBadRequest, response.Code)
	t.Require().NoError(err)
}

func (t *Test) Test_Signup_Failed_Birth_Date_Empty() {
	request := entity.UserSignupRequest{
		Email:     "user@email.com",
		Password:  "password",
		Name:      "User",
		BirthDate: "",
	}
	response, err := t.executePost(signupURL, request)
	t.Require().Equal(http.StatusBadRequest, response.Code)
	t.Require().NoError(err)
}

func (t *Test) Test_Signup_Failed_Invalid_Birth_Date_Format() {
	request := entity.UserSignupRequest{
		Email:     "user@email.com",
		Password:  "password",
		Name:      "User",
		BirthDate: "01-01-2024",
	}
	response, err := t.executePost(signupURL, request)
	t.Require().Equal(http.StatusBadRequest, response.Code)
	t.Require().NoError(err)
}

func (t *Test) Test_Signup_Failed_Gender_Empty() {
	request := entity.UserSignupRequest{
		Email:     "user@email.com",
		Password:  "password",
		Name:      "User",
		BirthDate: "2024-01-01",
		Gender:    "",
	}
	response, err := t.executePost(signupURL, request)
	t.Require().Equal(http.StatusBadRequest, response.Code)
	t.Require().NoError(err)
}

func (t *Test) Test_Signup_Failed_Invalid_Gender() {
	request := entity.UserSignupRequest{
		Email:     "user@email.com",
		Password:  "password",
		Name:      "User",
		BirthDate: "2024-01-01",
		Gender:    "GENDER",
	}
	response, err := t.executePost(signupURL, request)
	t.Require().Equal(http.StatusBadRequest, response.Code)
	t.Require().NoError(err)
}

func (t *Test) Test_Signup_Failed_Location_Empty() {
	request := entity.UserSignupRequest{
		Email:     "user@email.com",
		Password:  "password",
		Name:      "User",
		BirthDate: "2024-01-01",
		Gender:    "MALE",
		Location:  "",
	}
	response, err := t.executePost(signupURL, request)
	t.Require().Equal(http.StatusBadRequest, response.Code)
	t.Require().NoError(err)
}

func (t *Test) Test_Signup_Failed_Email_Already_Exists() {
	randomString := util.RandomString(6)
	email := randomString + "@email.com"
	request := entity.UserSignupRequest{
		Email:     email,
		Password:  "password",
		Name:      "Integration Test",
		BirthDate: time.Now().Format("2006-01-02"),
		Gender:    "MALE",
		Location:  "Indonesia",
	}
	response, err := t.executePost(signupURL, request)
	t.Require().Equal(http.StatusCreated, response.Code)
	t.Require().NoError(err)

	response, err = t.executePost(signupURL, request)
	t.Require().Equal(http.StatusConflict, response.Code)
	t.Require().NoError(err)
}

func (t *Test) Test_Signup_Success() {
	randomString := util.RandomString(6)
	email := randomString + "@email.com"
	request := entity.UserSignupRequest{
		Email:     email,
		Password:  "password",
		Name:      "Integration Test",
		BirthDate: time.Now().Format("2006-01-02"),
		Gender:    "MALE",
		Location:  "Indonesia",
	}
	response, err := t.executePost(signupURL, request)
	t.Require().Equal(http.StatusCreated, response.Code)
	t.Require().NoError(err)
}

func (t *Test) Test_Login_Failed_Bad_Request() {
	response, err := t.executePost(loginURL, nil)
	t.Require().Equal(http.StatusBadRequest, response.Code)
	t.Require().NoError(err)
}

func (t *Test) Test_Login_Failed_Email_Empty() {
	request := entity.UserSignupRequest{
		Email: "",
	}
	response, err := t.executePost(loginURL, request)
	t.Require().Equal(http.StatusBadRequest, response.Code)
	t.Require().NoError(err)
}

func (t *Test) Test_Login_Failed_Invalid_Email_Format() {
	request := entity.UserSignupRequest{
		Email: "invalid@email",
	}
	response, err := t.executePost(loginURL, request)
	t.Require().Equal(http.StatusBadRequest, response.Code)
	t.Require().NoError(err)
}

func (t *Test) Test_Login_Failed_Password_Empty() {
	request := entity.UserSignupRequest{
		Email:    "user@email.com",
		Password: "",
	}
	response, err := t.executePost(loginURL, request)
	t.Require().Equal(http.StatusBadRequest, response.Code)
	t.Require().NoError(err)
}

func (t *Test) Test_Login_Failed_User_Not_Found() {
	randomString := util.RandomString(6)
	email := randomString + "@email.com"
	request := entity.UserSignupRequest{
		Email:    email,
		Password: "password",
	}
	response, err := t.executePost(loginURL, request)
	t.Require().Equal(http.StatusNotFound, response.Code)
	t.Require().NoError(err)
}

func (t *Test) Test_Login_Failed_Invalid_Password() {
	randomString := util.RandomString(6)
	email := randomString + "@email.com"
	signUpReq := entity.UserSignupRequest{
		Email:     email,
		Password:  "password",
		Name:      "Integration Test",
		BirthDate: time.Now().Format("2006-01-02"),
		Gender:    "MALE",
		Location:  "Indonesia",
	}
	response, err := t.executePost(signupURL, signUpReq)
	t.Require().Equal(http.StatusCreated, response.Code)
	t.Require().NoError(err)

	loginReq := entity.UserLoginRequest{
		Email:    signUpReq.Email,
		Password: "passwodr",
	}
	response, err = t.executePost(loginURL, loginReq)
	t.Require().Equal(http.StatusBadRequest, response.Code)
	t.Require().NoError(err)
}

func (t *Test) Test_Login_Success() {
	randomString := util.RandomString(6)
	email := randomString + "@email.com"
	signUpReq := entity.UserSignupRequest{
		Email:     email,
		Password:  "password",
		Name:      "Integration Test",
		BirthDate: time.Now().Format("2006-01-02"),
		Gender:    "MALE",
		Location:  "Indonesia",
	}
	response, err := t.executePost(signupURL, signUpReq)
	t.Require().Equal(http.StatusCreated, response.Code)
	t.Require().NoError(err)

	loginReq := entity.UserLoginRequest{
		Email:    signUpReq.Email,
		Password: signUpReq.Password,
	}
	response, err = t.executePost(loginURL, loginReq)
	t.Require().Equal(http.StatusOK, response.Code)
	t.Require().NoError(err)
}
