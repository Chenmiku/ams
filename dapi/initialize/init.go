package initialize

import (
	"ams/dapi/config"
	"ams/dapi/httpserver"
	"util/runtime"
)

// func initialize(ctx context.Context) {
// 	mgo.Start(ctx)
// }

func Start(p *config.ProjectConfig) {
	runtime.MaxProc()
	server = httpserver.NewProjectHttpServer(p)
	//initialize(ctx)
}

func Wait() {
	defer beforeExit()
	server.Wait()
}
