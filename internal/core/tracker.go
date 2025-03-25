package core

import (
	"net/http"
	"time"
)

func TrackHandler(name string, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := &responseRecorder{ResponseWriter: w, statusCode: 200}
		h(rec, r)
		duration := time.Since(start)
		hasError := rec.statusCode >= 400
		TrackRequest(duration, hasError)
	}
}

type responseRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rr *responseRecorder) WriteHeader(code int) {
	rr.statusCode = code
	rr.ResponseWriter.WriteHeader(code)
}
