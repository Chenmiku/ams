package initialize

import (
	"ams/dapi/config"
	"ams/dapi/httpserver"
	"util/runtime"
)

func Start(p *config.ProjectConfig) {
	runtime.MaxProc()
	server = httpserver.NewProjectHttpServer(p)
}

func Wait() {
	defer beforeExit()
	server.Wait()
}
