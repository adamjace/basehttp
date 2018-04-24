package basehttp

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

// httpError is a helper to tidy up setting a HTTP error for a JSON response
func httpError(w http.ResponseWriter, statusCode int, err error) {

	setHeaderJSON(w, statusCode)

	if err == nil {
		return
	}

	response := Response{
		Error: err.Error(),
	}

	json.NewEncoder(w).Encode(response)
}

// httpData is a helper to tidy up setting a HTTP OK reponse with data
func httpData(w http.ResponseWriter, statusCode int, data interface{}) {

	setHeaderJSON(w, statusCode)

	response := Response{
		Data: data,
	}

	json.NewEncoder(w).Encode(response)
}

func setHeader(w http.ResponseWriter, contentType string, statusCode int) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
}

func setHeaderJSON(w http.ResponseWriter, statusCode int) {
	setHeader(w, "text/json", statusCode)
}
