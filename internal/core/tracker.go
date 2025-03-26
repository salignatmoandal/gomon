package core

import (
	"net/http"
	"time"
)

// -- TrackHandler Function --//
// The TrackHandler function wraps an existing HTTP handler, adding functionality to measure how long the handler takes to process a request and to determine if the response contains an error.
func TrackHandler(name string, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := &responseRecorder{ResponseWriter: w, statusCode: 200} // A custom response is created, which embeds the original http.ResponseWriter. It also initializes the statusCode to 200.
		h(rec, r)                                                    // Invokes the original handler with the custom response and the request.
		duration := time.Since(start)
		hasError := rec.statusCode >= 400
		TrackRequest(duration, hasError)
	}
}

// -- responseRecorder Struct --//
// The responseRecorder struct embeds the original http.ResponseWriter and adds a statusCode field.
type responseRecorder struct {
	http.ResponseWriter
	statusCode int
}

// The WriteHeader method of the responseRecorder struct overrides the default WriteHeader method of the http.ResponseWriter.
// It sets the statusCode field and then calls the original WriteHeader method.
func (rr *responseRecorder) WriteHeader(code int) {
	rr.statusCode = code
	rr.ResponseWriter.WriteHeader(code)
}
