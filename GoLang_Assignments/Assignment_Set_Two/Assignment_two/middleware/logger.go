package middleware

import (
	"log"
	"net/http"
	"time"
)


func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		
		log.Printf("Started %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)

		
		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(lrw, r)

		
		duration := time.Since(startTime)
		log.Printf("Completed %d %s in %v", lrw.statusCode, http.StatusText(lrw.statusCode), duration)
	})
}


type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}


func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
