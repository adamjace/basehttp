package basehttp

import (
	"net/http"

	"github.com/gorilla/mux"
)

type greetingController struct {
	baseController
}

func newGreetingController(c Config, s Service) greetingController {
	return greetingController{
		baseController: newBaseController(c, s),
	}
}

func (c greetingController) registerHandlers(r *mux.Router) {
	r.HandleFunc("/greet/{name}", c.handle).Methods("GET")
}

func (c greetingController) handle(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	name := vars["name"]

	greeting, err := c.service.Greet(name)
	if err != nil {
		httpError(w, http.StatusInternalServerError, err)
		return
	}

	httpData(w, http.StatusOK, greeting)
}
