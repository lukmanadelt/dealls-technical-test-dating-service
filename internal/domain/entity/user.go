package entity

import (
	"dealls-technical-test-dating-service/pkg/util"
	"errors"
	"slices"
	"strings"
	"time"
)

var genders = []string{"MALE", "FEMALE", "OTHER"}

// User is a struct that represents user attributes.
type User struct {
	ID                int       `json:"id"`
	Email             string    `json:"email"`
	Password          string    `json:"password"`
	Name              string    `json:"name"`
	BirthDate         time.Time `json:"birth_date"`
	Gender            string    `json:"gender"`
	Location          string    `json:"location"`
	ProfilePictureURL string    `json:"profile_picture_url"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// NewUser is a function used to initialize the user struct.
func NewUser(email, password, name, birthDate, gender, location, profilePictureURL string, createdAt, updatedAt time.Time) *User {
	birthDateParsed, _ := time.Parse("2006-01-02", birthDate)

	return &User{
		Email:             email,
		Password:          password,
		Name:              name,
		BirthDate:         birthDateParsed,
		Gender:            gender,
		Location:          location,
		ProfilePictureURL: profilePictureURL,
		CreatedAt:         createdAt,
		UpdatedAt:         updatedAt,
	}
}

// UserSignupRequest is a struct that represents user signup request body.
type UserSignupRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	BirthDate string `json:"birth_date"`
	Gender    string `json:"gender"`
	Location  string `json:"location"`
}

// Validate is a method for validating the attributes in the user signup request body.
func (u *UserSignupRequest) Validate() error {
	if u.Email == "" {
		return errors.New("email is required")
	}

	if !util.IsValidEmail(u.Email) {
		return errors.New("invalid email format")
	}

	if u.Password == "" {
		return errors.New("password is required")
	}

	if len(u.Password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	if u.Name == "" {
		return errors.New("name is required")
	}

	if u.BirthDate == "" {
		return errors.New("birth_date is required")
	}

	_, err := time.Parse("2006-01-02", u.BirthDate)
	if err != nil {
		return errors.New("invalid birth_date format. Format must be YYYY-MM-DD")
	}

	if u.Gender == "" || !slices.Contains(genders, strings.ToUpper(u.Gender)) {
		return errors.New("gender must be \"MALE\", \"FEMALE\", or \"OTHER\"")
	}

	if u.Location == "" {
		return errors.New("location is required")
	}

	return nil
}

// UserLoginRequest is a struct that represents user login request body.
type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Validate is a method for validating the attributes in the user login request body.
func (u *UserLoginRequest) Validate() error {
	if u.Email == "" {
		return errors.New("email is required")
	}

	if !util.IsValidEmail(u.Email) {
		return errors.New("invalid email format")
	}

	if u.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

// UserLoginResponse is a struct that represents user login response body.
type UserLoginResponse struct {
	Token string `json:"token"`
}

// NewUserLoginResponse is a function used to initialize the user login response struct.
func NewUserLoginResponse(token string) *UserLoginResponse {
	return &UserLoginResponse{
		Token: token,
	}
}
