package main

import (
	"gomon/internal/api"
	"gomon/internal/config"
	"gomon/internal/core"
)

func main() {
	conf := config.Load()

	core.LogInfo("Starting profiler on " + conf.ProfilePort)
	core.StartProfiler(conf.ProfilePort)

	core.LogInfo("Starting Gomon HTTP API on " + conf.ServerPort)
	api.StartServer(conf.ServerPort)
}
