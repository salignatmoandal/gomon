package core

import (
	"runtime"
	"sync"
	"time"
)

// -- Metrics Struct --//
// The Metrics struct holds the data for monitoring the application's performance.
// Field : mu is a sync.RWMutex used to synchronize access to the metrics.
// Field : RequestCount is the number of requests processed.
// Field : TotalLatency is the total time taken to process all requests.
// Field : ErrorCount is the number of requests that resulted in an error.
// Field : LastRequestTime is the time of the last request.
type Metrics struct {
	mu              sync.RWMutex
	RequestCount    int64
	TotalLatency    time.Duration
	ErrorCount      int64
	LastRequestTime time.Time
}

var metrics *Metrics

// -- init Function --//
// The init function initializes the metrics struct and sets the LastRequestTime to the current time.
func init() {
	metrics = &Metrics{
		LastRequestTime: time.Now(),
	}
}

func TrackRequest(duration time.Duration, hasError bool) {
	metrics.mu.Lock()
	defer metrics.mu.Unlock()

	metrics.RequestCount++
	metrics.TotalLatency += duration
	if hasError {
		metrics.ErrorCount++
	}
	metrics.LastRequestTime = time.Now()

}

func GetStats() map[string]interface{} {
	metrics.mu.RLock()
	defer metrics.mu.RUnlock()

	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	return map[string]interface{}{
		"request_count":     metrics.RequestCount,
		"avg_latency":       avgLatency(),
		"error_count":       metrics.ErrorCount,
		"last_request_time": metrics.LastRequestTime.Format(time.RFC3339),
		"goroutines":        runtime.NumGoroutine(),
		"memory_usage":      mem.TotalAlloc / 1024,
	}
}

func avgLatency() float64 {
	if metrics.RequestCount == 0 {
		return 0
	}
	return float64(metrics.TotalLatency.Milliseconds()) / float64(metrics.RequestCount)
}
