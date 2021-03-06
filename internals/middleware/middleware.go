package middleware

import (
	"log"
	"net/http"
	"time"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{ResponseWriter: w}
}

func (lrw *loggingResponseWriter) Status() int {
	return lrw.status
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	if lrw.wroteHeader {
		return
	}
	lrw.status = code
	lrw.ResponseWriter.WriteHeader(code)
	lrw.wroteHeader = true
}

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		lrw := newLoggingResponseWriter(w)
		next.ServeHTTP(lrw, r)
		log.Printf("[Incoming Request] Status: %d, Method: %s Path: %s Duration: %v ms", lrw.Status(), r.Method, r.URL.EscapedPath(), time.Since(start).Microseconds())
	})
}
