package basehttp

import "github.com/gorilla/mux"

type controller interface {
	registerHandlers(*mux.Router)
}

type baseController struct {
	config  Config
	service Service
}

func newBaseController(c Config, s Service) baseController {
	return baseController{
		config:  c,
		service: s,
	}
}
