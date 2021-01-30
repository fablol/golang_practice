package main

import (
	"fmt"
	"golang_practice/zinx/utils"
	"golang_practice/zinx/ziface"
	"golang_practice/zinx/znet"
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
	// fmt.Println("test pre handle")
	// _, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping"))
	// if err != nil {
	// 	fmt.Println("test before ping error")
	// }
}

// test  handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("test handle")
	// read data
	fmt.Printf("recv msgID: %d data : %s \n", request.GetMsgID(), request.GetData())
	err := request.GetConnection().SendMsg(1, []byte("123...456...789"))
	if err != nil {
		fmt.Println("handle error : ", err)
	}
}

// test post handle
func (this *PingRouter) PostHandle(request ziface.IRequest) {
	// fmt.Println("test post handle")
	// _, err := request.GetConnection().GetTCPConnection().Write([]byte("post ping"))
	// if err != nil {
	// 	fmt.Println("test post ping error")
	// }
}

func main() {
	// create server handle
	s := znet.NewServer(utils.GlobalObject.Name)
	// add router
	s.AddRouter(&PingRouter{})
	// start server
	s.Serve()
}
