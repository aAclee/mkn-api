package http

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

// Logger standardization
type Logger interface {
	// Info will log arguments in the manner of fmt.Print.
	Info(v ...interface{})

	// Infof will log arguments are handled in the manner of fmt.Printf.
	Infof(format string, v ...interface{})

	// Error will log arguments in the manner of fmt.Print.
	Error(v ...interface{})

	// Errorf will log arguments are handled in the manner of fmt.Printf.
	Errorf(format string, v ...interface{})
}

// Config contains all the required configurations to be passed into server creation
type Config struct {
	Log  Logger
	Port int
}

// Server represents a http server
type Server struct {
	listener net.Listener
	log      Logger
	port     int
}

// ListenAndServe listens on the TCP network and serves handler to handle requests
func (s *Server) ListenAndServe() error {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return err
	}

	s.listener = ln

	s.log.Infof("Starting server on port %d...\n", s.port)
	return http.Serve(s.listener, s.CreateMux())
}

// Close stops the listener if it's open
func (s *Server) Close() {
	if s.listener != nil {
		s.listener.Close()
	}
}

// CreateMux sets up api routes and catch all
func (s *Server) CreateMux() *mux.Router {
	r := mux.NewRouter()

	return r
}

// CreateServer creates a server ready for listening
func CreateServer(config Config) Server {
	return Server{
		log:  config.Log,
		port: config.Port,
	}
}
