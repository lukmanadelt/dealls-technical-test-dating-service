// Package integration_test contains setup code for your integration tests, such as initializing the database, starting the server, and the actual integration test cases.
package integration_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"sync"

	"dealls-technical-test-dating-service/internal/domain/service"
	"dealls-technical-test-dating-service/internal/domain/usecase"
	"dealls-technical-test-dating-service/internal/infrastructure/auth"
	"dealls-technical-test-dating-service/internal/infrastructure/config"
	"dealls-technical-test-dating-service/internal/infrastructure/database"
	"dealls-technical-test-dating-service/internal/infrastructure/repository"
	"dealls-technical-test-dating-service/internal/interface/controller"
	"dealls-technical-test-dating-service/internal/interface/routes"
	"dealls-technical-test-dating-service/pkg/util"

	"github.com/caarlos0/env"
	"github.com/emicklei/go-restful/v3"
	"github.com/joho/godotenv"
	"github.com/parnurzeal/gorequest"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

// Constants.
const (
	BasePath            = "/e-wallet"
	AuthorizationHeader = "Authorization"
)

var once sync.Once

// HTTPTestCaller is a struct that represents an HTTP call request.
type HTTPTestCaller struct {
	handler  http.Handler
	request  *http.Request
	response interface{}
}

func call(handler http.Handler) *HTTPTestCaller {
	return &HTTPTestCaller{
		handler: handler,
	}
}

func (h *HTTPTestCaller) to(request *http.Request, _ ...interface{}) *HTTPTestCaller {
	h.request = request

	return h
}

func (h *HTTPTestCaller) execute() (*httptest.ResponseRecorder, interface{}, error) {
	recorder := httptest.NewRecorder()
	h.handler.ServeHTTP(recorder, h.request)

	var err error
	if h.response != nil {
		err = json.Unmarshal(recorder.Body.Bytes(), &h.response)
	}

	return recorder, h.response, err
}

// Test is a struct that represents the integration tests suite.
type Test struct {
	suite.Suite
	container *restful.Container
}

// SetupSuite is a method for setup the integration tests suite.
func (t *Test) SetupSuite() {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true)

	cfg, err := loadConfig()
	if err != nil {
		logrus.Fatalf("Failed to load configuration: %s", err.Error())
	}
	t.Require().NoError(err)
	t.Require().NotNil(cfg)

	postgres, err := database.NewPostgres(cfg)
	if err != nil {
		logrus.Fatalf("Failed to initialize PostgreSQL: %s", err.Error())
	}
	t.Require().NoError(err)
	t.Require().NotNil(postgres)

	err = postgres.Migrate(cfg.PostgresDBName)
	if err != nil {
		logrus.Fatalf("Failed to migrate PostgreSQL: %s", err.Error())
	}
	t.Require().NoError(err)

	container := restful.NewContainer()

	userRepo := repository.NewUserRepository(postgres.Client)
	userService := service.NewUserService(userRepo)

	profileRepo := repository.NewProfileRepository(postgres.Client)
	profileService := service.NewProfileService(profileRepo)

	jwt := auth.NewJWTClaims(cfg.JWTExpiration)
	userUsecase := usecase.NewUserUsecase(userService, profileService, cfg, jwt, util.HashPassword, util.IsValidPasswordHash)
	userController := controller.NewUserController(userUsecase)
	routes.RegisterUserRoutes(container, cfg.BasePath, userController)

	t.container = container
}

func (t *Test) execute(request *http.Request, _ ...interface{}) (*httptest.ResponseRecorder, interface{}, error) {
	return call(t.container).to(request).execute()
}

func (t *Test) executePost(url string, request interface{}) (*httptest.ResponseRecorder, error) {
	response, _, err := t.execute(gorequest.New().
		Post(url).
		Type(restful.MIME_JSON).
		Send(request).
		MakeRequest())

	return response, err
}

func loadConfig() (*config.Config, error) {
	var err error

	once.Do(func() {
		var filePath string

		filePath, err = util.GetFile(".env")
		if err == nil {
			err = godotenv.Load(filePath)
			if err != nil {
				return
			}
		}
	})

	if err != nil {
		return nil, err
	}

	var cfg config.Config
	err = env.Parse(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
