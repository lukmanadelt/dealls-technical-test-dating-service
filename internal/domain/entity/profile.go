// Package entity holds the core entities (models) of the application.
package entity

import "time"

// Profile is a struct that represents profile attributes.
type Profile struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Bio       string    `json:"bio"`
	Interests string    `json:"interests"`
	Verified  bool      `json:"verified"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewProfile is a function used to initialize the profile struct.
func NewProfile(userID int, bio, interests string, verified bool, createdAt, updatedAt time.Time) *Profile {
	return &Profile{
		UserID:    userID,
		Bio:       bio,
		Interests: interests,
		Verified:  verified,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
