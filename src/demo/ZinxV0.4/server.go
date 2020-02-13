package main

import (
	"fmt"
	"zinx/utils"
	"zinx/ziface"
	"zinx/znet"
)

/*
	base on zinx
*/

// ping test
// custom router
type PingRouter struct {
	znet.BaseRouter
}

// test pre handle
func (this *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("test pre handle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping"))
	if err != nil {
		fmt.Println("test before ping error")
	}
}

// test  handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("test handle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping..."))
	if err != nil {
		fmt.Println("test ping error")
	}
}

// test post handle
func (this *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("test post handle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("post ping"))
	if err != nil {
		fmt.Println("test post ping error")
	}
}

func main() {
	// create server handle
	s := znet.NewServer(utils.GlobalObject.Name)
	// add router
	s.AddRouter(&PingRouter{})
	// start server
	s.Serve()
}
