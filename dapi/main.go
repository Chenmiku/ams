package main

import (
	"context"
	"ams/dapi/config"
	"ams/dapi/initialize"
	//"ams/dapi/o/auth/session"
	//"time"
	//"os/user"
)

// func clearSession() {
//     session.Delete(id)
// }
// func duration() {
//     t := time.Now()
//     n := time.Date(t.Year(), t.Month(), t.Day(), 12, 0, 0, 0, t.Location())
//     d := n.Sub(t)
//     if d < 0 {
//         n = n.Add(10 * time.Second)
//         d = n.Sub(t)
//     }
//     for {
//         time.Sleep(d)
//         d = 10 * time.Second
//         clearSession()
//     }
// }

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	initialize.Start(ctx, config.ReadConfig())
	initialize.Wait()
	// auto clear session
	//duration()
}