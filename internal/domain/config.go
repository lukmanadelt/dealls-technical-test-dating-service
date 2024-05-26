package domain

// Config is an interface that represents the configuration requirements of the domain.
type Config interface {
	GetJWTKey() string
}
