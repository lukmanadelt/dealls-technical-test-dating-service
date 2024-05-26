// Package main initializes the application and invokes the necessary components to start.
package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	"dealls-technical-test-dating-service/internal/domain/service"
	"dealls-technical-test-dating-service/internal/domain/usecase"
	"dealls-technical-test-dating-service/internal/infrastructure/auth"
	"dealls-technical-test-dating-service/internal/infrastructure/config"
	"dealls-technical-test-dating-service/internal/infrastructure/database"
	"dealls-technical-test-dating-service/internal/infrastructure/log"
	"dealls-technical-test-dating-service/internal/infrastructure/repository"
	"dealls-technical-test-dating-service/internal/infrastructure/server"
	"dealls-technical-test-dating-service/internal/interface/controller"
	"dealls-technical-test-dating-service/internal/interface/routes"
	"dealls-technical-test-dating-service/pkg/util"

	"github.com/sirupsen/logrus"
)

func main() {
	log.InitLog()

	cfg, err := config.LoadConfig()
	if err != nil {
		for _, help := range cfg.Help() {
			fmt.Println(help)
		}

		logrus.Fatalf("Failed to load configuration: %s", err.Error())
	}

	logLevel, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		logrus.Fatalf("Failed to parse log level: %s", err.Error())
	}
	logrus.SetLevel(logLevel)

	postgres, err := database.NewPostgres(cfg)
	if err != nil {
		logrus.Fatalf("Failed to initialize PostgreSQL: %s", err.Error())
	}

	err = postgres.Migrate(cfg.PostgresDBName)
	if err != nil {
		logrus.Fatalf("Failed to migrate PostgreSQL: %s", err.Error())
	}

	server := server.NewServer(cfg.Port)
	if server == nil {
		logrus.Fatal("Failed to initialize server")
	}

	userRepo := repository.NewUserRepository(postgres.Client)
	userService := service.NewUserService(userRepo)

	profileRepo := repository.NewProfileRepository(postgres.Client)
	profileService := service.NewProfileService(profileRepo)

	jwt := auth.NewJWTClaims(cfg.JWTExpiration)
	userUsecase := usecase.NewUserUsecase(userService, profileService, cfg, jwt, util.HashPassword, util.IsValidPasswordHash)
	userController := controller.NewUserController(userUsecase)
	routes.RegisterUserRoutes(server.Container, cfg.BasePath, userController)

	defer func() {
		r := recover()
		if r == nil {
			return
		}

		fmt.Printf("Panic: %+v\n", r)
		debug.PrintStack()
		os.Exit(1)
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		s := <-sig
		logrus.Infof("Received quit signal %+v", s)
		server.Stop()
	}()

	server.Start()
}
