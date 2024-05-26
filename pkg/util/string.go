package util

import (
	"math/rand"
	"regexp"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// IsValidEmail is a function to validate an email.
func IsValidEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	return regexp.MustCompile(regex).MatchString(email)
}

// RandomString is a function to generate random strings.
func RandomString(length int) string {
	bytes := make([]byte, length)
	for i := range bytes {
		bytes[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(bytes)
}
