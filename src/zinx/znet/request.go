package znet

import (
	"zinx/ziface"
)

// Request inface
type Request struct {
	// connection
	conn ziface.IConnection
	// data
	msg ziface.IMessage
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.msg.GetMsgData()
}

func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgID()
}
