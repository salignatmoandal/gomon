package api

import (
	"encoding/json"
	"fmt"
	"gomon/internal/core"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Gomon is healthy\n"))
}

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	stats := core.GetStats()
	if err := json.NewEncoder(w).Encode(stats); err != nil {
		http.Error(w, "Erreur interne", http.StatusInternalServerError)
		return
	}
}

func PrometheusHandler(w http.ResponseWriter, r *http.Request) {
	stats := core.GetStats()
	w.Header().Set("Content-Type", "text/plain")

	for k, v := range stats {
		w.Write([]byte(fmt.Sprintf("gomon_%s{version=\"%s\"} %s\n",
			k,
			"0.1.0",
			formatPrometheusValue(v))))
	}
}

func formatPrometheusValue(v interface{}) string {
	switch val := v.(type) {
	case int:
		return fmt.Sprintf("%d", val)
	case float64:
		return fmt.Sprintf("%.3f", val)
	default:
		return fmt.Sprintf("\"%v\"", val)
	}
}
