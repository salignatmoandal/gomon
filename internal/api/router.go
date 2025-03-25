package api

import (
	"net/http"
)

func StartServer(port string) {
	http.HandleFunc("/health", HealthHandler)
	http.HandleFunc("/stats", StatsHandler)
	http.HandleFunc("/metrics", PrometheusHandler)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic("failed to start HTTP server: " + err.Error())
	}
}
