package znet

import (
	// "fmt"
	// "net"
	"zinx/ziface"
)

type BaseRouter struct{}

// before connection
func (br *BaseRouter) PreHandle(request ziface.IRequest) {}

// handleing connection
func (br *BaseRouter) Handle(request ziface.IRequest) {}

// after connection
func (br *BaseRouter) PostHandle(request ziface.IRequest) {}
