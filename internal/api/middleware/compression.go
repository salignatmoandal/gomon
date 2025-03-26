package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

// gzipResponseWriter wraps an http.ResponseWriter and an io.Writer.
// It is used to intercept writes and compress them using gzip.
type gzipResponseWriter struct {
	io.Writer           // The gzip writer that compresses the output.
	http.ResponseWriter // The original response writer.
}

// Write overrides the default Write method.
// It writes the data to the underlying gzip writer.
func (g gzipResponseWriter) Write(b []byte) (int, error) {
	return g.Writer.Write(b)
}

// CompressHandler is a middleware that compresses HTTP responses using gzip.
// It checks if the client supports gzip encoding and wraps the response accordingly.
func CompressHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check if the client accepts gzip encoding by inspecting the "Accept-Encoding" header.
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			// If gzip is not supported, call the next handler without modifying the response.
			next(w, r)
			return
		}

		// Initialize a new gzip writer that wraps the original ResponseWriter.
		gz := gzip.NewWriter(w)
		// Ensure the gzip writer is closed after the response is written.
		defer gz.Close()

		// Set the appropriate HTTP headers to indicate the response is gzip-compressed.
		w.Header().Set("Content-Encoding", "gzip")
		// The "Vary" header informs caches that the response varies based on the Accept-Encoding header.
		w.Header().Set("Vary", "Accept-Encoding")

		// Create a new gzipResponseWriter that uses the gzip writer.
		gzw := gzipResponseWriter{
			Writer:         gz,
			ResponseWriter: w,
		}

		// Call the next handler using the custom gzipResponseWriter.
		next(gzw, r)
	}
}
