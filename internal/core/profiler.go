package core

import (
	"net/http"

	_ "net/http/pprof"
)

// StartProfiler starts the profiler on the given port
func StartProfiler(port string) {
	go func() {
		LogInfo("Profiler started on port " + port)
		http.ListenAndServe(":"+port, nil)
	}()

}
