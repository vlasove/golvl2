package apiserver

import "net/http"

// responseWriter ...
type responseWriter struct {
	http.ResponseWriter
	code int
}

// WriteHeader ...
func (w *responseWriter) WriteHeader(statusCode int) {
	w.code = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
