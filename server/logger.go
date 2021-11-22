package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Logger(r *mux.Router) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			start := time.Now()
			defer func() {
				log.Printf(
					"[%s] %s %s %s (%s)",
					req.Method,
					req.Host,
					req.URL.Path,
					req.URL.RawQuery,
					time.Since(start),
				)
			}()
			next.ServeHTTP(w, req)
		})
	}
}
