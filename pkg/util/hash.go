package util

import "golang.org/x/crypto/bcrypt"

// HashPassword is a function to hash a password.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(bytes), err
}

// IsValidPasswordHash is a function to validate a password hash.
func IsValidPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}
