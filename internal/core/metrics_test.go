package core

import (
	"testing"
	"time"
)

func TestTrackRequestAndGetStats(t *testing.T) {
	// Reset metrics manually for test
	metrics = &Metrics{}

	// Simulate two requests
	TrackRequest(100*time.Millisecond, false)
	TrackRequest(200*time.Millisecond, true)

	stats := GetStats()

	// Check request count
	if stats["request_count"].(int64) != 2 {
		t.Errorf("Expected 2 requests, got %v", stats["request_count"])
	}

	// Check error count
	if stats["error_count"].(int64) != 1 {
		t.Errorf("Expected 1 error, got %v", stats["error_count"])
	}

	// Check average latency
	expectedAvg := float64((100 + 200)) / 2.0
	if stats["avg_latency"].(float64) != expectedAvg {
		t.Errorf("Expected avg latency %v ms, got %v", expectedAvg, stats["avg_latency"])
	}

	// Check goroutines
	if stats["goroutines"].(int) <= 0 {
		t.Errorf("Expected some goroutines, got %v", stats["goroutines"])
	}

	// Check memory usage is not zero (could be flaky on some archs)
	if stats["memory_usage"].(uint64) == 0 {
		t.Errorf("Expected non-zero memory usage, got %v", stats["memory_usage"])
	}

	// Check LastRequestTime is recent (within 2 seconds)
	lastReqStr := stats["last_request_time"].(string)
	lastReqTime, err := time.Parse(time.RFC3339, lastReqStr)
	if err != nil {
		t.Fatalf("Invalid time format: %v", lastReqStr)
	}
	if time.Since(lastReqTime) > 2*time.Second {
		t.Errorf("Last request time is too old: %v", lastReqTime)
	}
}
