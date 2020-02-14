package ziface

import (
	"net"
)

type IConnection interface {
	Start()
	Stop()
	GetTCPConnection() *net.TCPConn
	GetConnID() uint32
	GetRemoteAddr() net.Addr
	SendMsg(uint32, []byte) error
}

type HandleFunc func(*net.TCPConn, []byte, int) error
