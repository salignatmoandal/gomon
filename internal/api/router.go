package api

import (
	"gomon/internal/api/middleware"
	"net/http"
)

func StartServer(port string) {
	http.HandleFunc("/health", HealthHandler)
	http.HandleFunc("/stats", middleware.CompressHandler(StatsHandler))
	http.HandleFunc("/metrics", middleware.CompressHandler(PrometheusHandler))

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic("failed to start HTTP server: " + err.Error())
	}
}
