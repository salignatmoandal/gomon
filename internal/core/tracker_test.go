package core

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestTrackHandler(t *testing.T) {
	// Reset metrics
	metrics = &Metrics{}

	handlerOK := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(50 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}

	handlerErr := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Wrap handlers with middleware
	trackedOK := TrackHandler("ok", handlerOK)
	trackedErr := TrackHandler("err", handlerErr)

	// Call the OK handler
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	trackedOK(w, req)

	// Call the Error handler
	req2 := httptest.NewRequest("GET", "/", nil)
	w2 := httptest.NewRecorder()
	trackedErr(w2, req2)

	// Validate metrics
	stats := GetStats()

	if stats["request_count"].(int64) != 2 {
		t.Errorf("Expected 2 requests, got %v", stats["request_count"])
	}

	if stats["error_count"].(int64) != 1 {
		t.Errorf("Expected 1 error, got %v", stats["error_count"])
	}

	avg := stats["avg_latency"].(float64)
	if avg < 50 || avg > 150 {
		t.Errorf("Unexpected average latency: %vms", avg)
	}
}
