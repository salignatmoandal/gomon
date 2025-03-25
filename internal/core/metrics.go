package core

import (
	"runtime"
	"sync"
	"time"
)

type Metrics struct {
	mu              sync.RWMutex
	RequestCount    int64
	TotalLatency    time.Duration
	ErrorCount      int64
	LastRequestTime time.Time
}

func (m *Metrics) IncRequestCount() {
	panic("unimplemented")
}

var metrics *Metrics

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
