package apiserver

import (
	"fmt"
	"net/http"
	"time"
)

func (s *APIServer) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.logger.Println(
			fmt.Sprintf("started %s %s", r.Method, r.RequestURI),
		)
		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}

		next.ServeHTTP(rw, r)

		s.logger.Println(
			fmt.Sprintf("completed with %d %s in %v",
				rw.code,
				http.StatusText(rw.code),
				time.Since(start),
			),
		)
	})
}
