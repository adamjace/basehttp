package basehttp

import (
	"log"
	"net/http"
	"time"
)

type middleware struct{}

func newMiddleware() middleware {
	return middleware{}
}

func (m middleware) Log(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		next.ServeHTTP(w, r)

		elapsed := float64(time.Since(start).Nanoseconds()) / float64(1000000)

		log.Printf("method=%s address=%s path=%s elapsed=%v",
			r.Method,
			r.RemoteAddr,
			r.RequestURI,
			elapsed)
	})

}
