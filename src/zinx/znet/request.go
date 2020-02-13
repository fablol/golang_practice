package znet

import (
	"zinx/ziface"
)

// Request inface
type Request struct {
	// connection
	conn ziface.IConnection
	// data
	data []byte
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}
