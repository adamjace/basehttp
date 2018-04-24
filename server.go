package basehttp

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Server provides HTTP server functionality
type Server struct {
	config  Config
	service Service
}

// NewServer creates a new Server
func NewServer(c Config) Server {
	s := NewService(c)

	return Server{
		config:  c,
		service: s,
	}
}

// Start kicks off the server, listening for all HTTP requests
func (s Server) Start() error {
	r := mux.NewRouter()

	r.HandleFunc("/ping", pingHandler)

	controllers := []controller{
		newGreetingController(s.config, s.service),
	}

	for _, c := range controllers {
		c.registerHandlers(r)
	}

	return http.ListenAndServe(s.config.Bind, r)
}
