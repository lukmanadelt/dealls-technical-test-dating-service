// Package config contains configuration-related code.
package config

import (
	"fmt"
	"reflect"
	"time"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

const (
	helpTitle  = "Environment variables config:"
	envName    = "env"
	envDefault = "envDefault"
	envDocs    = "envDocs"
)

// Config is a struct that represents a list of environment variables for configuration.
type Config struct {
	Port     string `env:"PORT"      envDefault:"8080"    envDocs:"The port that the service listens to"`
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"    envDocs:"Log level"`
	BasePath string `env:"BASE_PATH" envDefault:"/dating" envDocs:"Service API base path"`

	PostgresHost                  string        `env:"POSTGRES_HOST,required"           envDocs:"Host to connect to the PostgreSQL"`
	PostgresPort                  string        `env:"POSTGRES_PORT,required"           envDocs:"Port to connect to the PostgreSQL"`
	PostgresUser                  string        `env:"POSTGRES_USER,required"           envDocs:"User to connect to the PostgreSQL"`
	PostgresPassword              string        `env:"POSTGRES_PASSWORD,required"       envDocs:"Password to connect to the PostgreSQL"`
	PostgresDBName                string        `env:"POSTGRES_DB_NAME"                 envDocs:"PostgreSQL database name"`
	PostgresSSLMode               string        `env:"POSTGRES_SSL_MODE"                envDocs:"PostgreSQL SSL Mode"`
	PostgresMaxOpenConnections    int           `env:"POSTGRES_MAX_OPEN_CONNECTIONS"    envDefault:"30"                                 envDocs:"Maximum number of open connections to the PostgreSQL"`
	PostgresMaxIdleConnections    int           `env:"POSTGRES_MAX_IDLE_CONNECTIONS"    envDefault:"10"                                 envDocs:"Maximum number of buffered connections to the PostgreSQL"`
	PostgresConnectionMaxLifetime time.Duration `env:"POSTGRES_CONNECTION_MAX_LIFETIME" envDefault:"3600s"                              envDocs:"Maximum lifetime of an idle connection to PostgreSQL in seconds"`

	JWTKey        string        `env:"JWT_KEY"        envDocs:"Key for the JWT signed string"`
	JWTExpiration time.Duration `env:"JWT_EXPIRATION" envDefault:"24h"                        envDocs:"JWT expiration duration"`
}

// LoadConfig is the function used to load the configuration..
func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = env.Parse(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

// Help is a method for displaying detailed information on each configuration.
func (c Config) Help() []string {
	config := reflect.TypeOf(c)

	help := make([]string, 1+config.NumField())
	help[0] = helpTitle

	for i := 0; i < config.NumField(); i++ {
		field := config.Field(i)
		env := field.Tag.Get(envName)
		envDefault := field.Tag.Get(envDefault)
		envDocs := field.Tag.Get(envDocs)
		help[i+1] = fmt.Sprintf("%s | Description: %s | Default: %s", env, envDocs, envDefault)
	}

	return help
}

// GetJWTKey is a method for getting the JWT key.
func (c Config) GetJWTKey() string {
	return c.JWTKey
}
