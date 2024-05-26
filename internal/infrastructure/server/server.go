// Package server contains server connection and setup.
package server

import (
	"net"
	"net/http"

	"github.com/emicklei/go-restful/v3"
	"github.com/sirupsen/logrus"
)

// Server is a struct that represents the server attributes.
type Server struct {
	port      string
	listener  net.Listener
	Container *restful.Container
}

// NewServer is a function used to initialize the server.
func NewServer(port string) *Server {
	listener, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		logrus.Errorf("Failed to listen on port %s: %s", port, err.Error())

		return nil
	}

	server := &Server{
		port:      port,
		listener:  listener,
		Container: restful.NewContainer(),
	}
	server.Container.Filter(server.Container.OPTIONSFilter)

	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders:  []string{},
		AllowedMethods: []string{http.MethodPost, http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodDelete},
		AllowedHeaders: []string{"Access-Control-Allow-Origin", "Access-Control-Allow-Methods", "Authorization", "Content-Type", "Accept"},
		CookiesAllowed: true,
		Container:      server.Container,
		MaxAge:         0,
	}
	server.Container.Filter(cors.Filter)

	return server
}

// Start is a method for starting the server.
func (s *Server) Start() {
	logrus.Infof("Starting the server on port %s...", s.port)

	err := http.Serve(s.listener, s.Container)
	if err != nil {
		logrus.Errorf("Failed to serve: %s", err.Error())
	}
}

// Stop is a method for stopping the server.
func (s *Server) Stop() {
	logrus.Infof("Stopping the server on port %s...", s.port)
	err := s.listener.Close()
	if err != nil {
		logrus.Errorf("Failed to close the listener: %s", err.Error())
	}
}
