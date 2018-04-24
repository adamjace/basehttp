package basehttp

import (
	"io"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	setHeader(w, "text/plain", http.StatusOK)
	io.WriteString(w, "pong")
}
