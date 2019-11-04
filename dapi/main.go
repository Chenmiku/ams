package main

import (
	"ams/dapi/config"
	"ams/dapi/initialize"
)

func main() {
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()
	initialize.Start(config.ReadConfig())
	initialize.Wait()
}
