package main

import (
	"zinx/znet"
)

/*
	base on zinx
*/

func main() {
	// create server handle
	s := znet.NewServer("zinx v0.1.1.2")
	// start server
	s.Serve()
}
