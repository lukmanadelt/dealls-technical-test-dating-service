// Package auth contains implementation of the auth interface defined in the domain package.
package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWTClaims is a struct that represents the attributes used to generate the JWT.
type JWTClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// NewJWTClaims is a function used to initialize the JWT claims.
func NewJWTClaims(expiration time.Duration) *JWTClaims {
	return &JWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expiration).Unix(),
		},
	}
}

// GenerateToken is a method for generating JWT.
func (j *JWTClaims) GenerateToken(email, key string) (string, error) {
	j.Email = email
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, j)

	return token.SignedString([]byte(key))
}
