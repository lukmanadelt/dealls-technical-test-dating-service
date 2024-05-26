// Package routes contains the routes for the web service.
package routes

import (
	"net/http"

	"dealls-technical-test-dating-service/internal/domain/entity"
	"dealls-technical-test-dating-service/internal/interface/controller"

	"github.com/emicklei/go-restful/v3"
)

// RegisterUserRoutes is a function to register routes for user APIs.
func RegisterUserRoutes(container *restful.Container, basePath string, controller *controller.UserController) {
	webService := new(restful.WebService).Path(basePath)
	webService.Route(webService.
		POST("/v1/users/signup").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON).
		Reads(entity.UserSignupRequest{}).
		Returns(http.StatusCreated, http.StatusText(http.StatusCreated), nil).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil).
		Returns(http.StatusConflict, http.StatusText(http.StatusConflict), nil).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), nil).
		To(controller.Signup))
	webService.Route(webService.
		POST("/v1/users/login").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON).
		Reads(entity.UserLoginRequest{}).
		Returns(http.StatusOK, http.StatusText(http.StatusOK), entity.UserLoginResponse{}).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil).
		Returns(http.StatusNotFound, http.StatusText(http.StatusNotFound), nil).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), nil).
		To(controller.Login))

	container.Add(webService)
}
