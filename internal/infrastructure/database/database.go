// Package database contains database connection and setup.
package database

import (
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	"dealls-technical-test-dating-service/internal/infrastructure/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	postgresMigrate "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // required by migration.
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	migrationDirectory = "/internal/infrastructure/database/migration"
	windowsOS          = "windows"
)

// Postgres is a struct that represents the PostgreSQL client.
type Postgres struct {
	Client *gorm.DB
}

// NewPostgres is a function used to initialize the PostgreSQL client.
func NewPostgres(cfg *config.Config) (*Postgres, error) {
	postgresDSN := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresDBName, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresSSLMode)
	postgresClient, err := gorm.Open(postgres.Open(postgresDSN), &gorm.Config{
		DisableAutomaticPing:   true,
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		return nil, err
	}

	postgresDB, err := postgresClient.DB()
	if err != nil {
		return nil, err
	}

	if cfg.PostgresMaxOpenConnections > 0 {
		postgresDB.SetMaxOpenConns(cfg.PostgresMaxOpenConnections)
	}

	if cfg.PostgresMaxIdleConnections > 0 {
		postgresDB.SetMaxIdleConns(cfg.PostgresMaxIdleConnections)
	}

	postgresDB.SetConnMaxLifetime(cfg.PostgresConnectionMaxLifetime)

	return &Postgres{
		Client: postgresClient,
	}, nil
}

func getMigrationDirectory(targetDirectory string) (string, error) {
	directory, err := os.Getwd()
	if err != nil {
		return "", err
	}

	directories := strings.Split(directory, string(os.PathSeparator))
	if runtime.GOOS == windowsOS && len(directories) > 0 {
		directories = directories[1:]
	}

	for i := len(directories); i > 0; i-- {
		directory := "/"
		for j := 0; j < i; j++ {
			directory = path.Join(directory, directories[j])
		}

		directory = path.Join(directory, targetDirectory)
		fileInfo, err := os.Stat(directory)
		if os.IsNotExist(err) {
			continue
		}

		if fileInfo.IsDir() {
			return directory, nil
		}
	}

	return "", fmt.Errorf("Directory %s not found", targetDirectory)
}

// Migrate is the method used to perform migration queries.
func (p *Postgres) Migrate(databaseName string) error {
	logrus.Info("Starting PostgreSQL migration...")

	migrationDirectory, err := getMigrationDirectory(migrationDirectory)
	if err != nil {
		return err
	}
	if migrationDirectory == "" {
		return errors.New("The migration directory is empty")
	}

	var driver database.Driver
	retry := 0
	timeout := time.Now().UTC().Add(time.Second * 10)
	for time.Now().UTC().Before(timeout) {
		postgresDB, err := p.Client.DB()
		if err != nil {
			return fmt.Errorf("Failed to get PostgreSQL DB: %s", err.Error())
		}

		driver, err = postgresMigrate.WithInstance(postgresDB, &postgresMigrate.Config{})
		if err == nil {
			break
		}

		retry++
		logrus.Errorf("Connection to database error: %s", err.Error())
		logrus.Infof("Retry connection to database: %d", retry)
		time.Sleep(time.Second)
	}

	if driver == nil {
		return errors.New("Failed to initiate database migration")
	}

	mig, err := migrate.NewWithDatabaseInstance("file://"+migrationDirectory, databaseName, driver)
	if err != nil {
		return err
	}

	err = mig.Up()
	if errors.Is(err, migrate.ErrNoChange) {
		logrus.Info("Migration was not performed because no changes were detected")

		return nil
	}
	if err != nil {
		return err
	}

	logrus.Info("PostgreSQL migration successful")

	return nil
}
