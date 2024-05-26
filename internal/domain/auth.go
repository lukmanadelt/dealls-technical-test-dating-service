// Package domain contains domain-specific logic and entities.
package domain

// Auth is an interface that represents the authentication functionality needed by the domain.
type Auth interface {
	GenerateToken(email, key string) (string, error)
}
