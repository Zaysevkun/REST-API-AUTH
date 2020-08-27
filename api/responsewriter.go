package api

import "net/http"

//response writer wrapper
type responseWriter struct {
	http.ResponseWriter
	code int
}

// writeHeader wrapper that allows us to get statusCode from responseWriter
func (w *responseWriter) WriteHeader(statusCode int) {
	w.code = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
