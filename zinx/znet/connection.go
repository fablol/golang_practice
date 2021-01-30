package znet

import (
	"errors"
	"fmt"
	"golang_practice/zinx/ziface"
	"io"
	"net"
)

type Connection struct {
	Conn     *net.TCPConn
	ConnID   uint32
	isClosed bool
	// handleAPI   ziface.HandleFunc
	ExitChannel chan bool
	// current router
	Router ziface.IRouter
}

// init
func NewConnection(conn *net.TCPConn, ConnID uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   ConnID,
		isClosed: false,
		Router:   router,
		// handleAPI:   callback_api,
		ExitChannel: make(chan bool, 1),
	}
	return c
}

func (c *Connection) StartReader() {
	fmt.Println("start read goroutine.. ConnID = ", c.ConnID)
	defer fmt.Println("ConnID = ", c.ConnID, " reader is exit")
	defer c.Stop()

	for {
		// read data
		dp := NewDataPack()
		headData := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(c.GetTCPConnection(), headData)
		if err != nil {
			fmt.Println("read msg head error : ", err)
			break
		}

		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack msg head error : ", err)
			break
		}

		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read msg data error : ", err)
				break
			}
		}

		msg.SetMsgData(data)
		// get conn request
		req := Request{
			conn: c,
			msg:  msg,
		}

		// use router
		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)

		// c.Router.PreHandle(req)
		// err = c.handleAPI(c.Conn, buf, count)

		// if err != nil {
		// 	fmt.Println("connection : ", c.ConnID, " handle err : ", err)
		// 	break
		// }
	}
}

func (c *Connection) StartWriter() {
	fmt.Println("start write goroutine.. ConnID = ", c.ConnID)
}

func (c *Connection) Start() {
	fmt.Println("connection start.. ConnID = ", c.ConnID)

	// read data
	go c.StartReader()
	// write data

}
func (c *Connection) Stop() {
	fmt.Println("connection stop.. ConnID = ", c.ConnID)

	//
	if c.isClosed == true {
		return
	}

	c.isClosed = true
	// close
	c.Conn.Close()

	close(c.ExitChannel)
}
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}
func (c *Connection) GetRemoteAddr() net.Addr {
	return nil
}

// send msg
func (c *Connection) SendMsg(msgID uint32, data []byte) error {
	if c.isClosed == true {
		return errors.New("connection is closed!")
	}

	// pack data
	dp := NewDataPack()
	msg, err := dp.Pack(NewMsgPackage(msgID, data))

	if err != nil {
		fmt.Println("pack msg error : ", msgID)
		return errors.New("package msg error!")
	}

	if _, err := c.Conn.Write(msg); err != nil {
		fmt.Println("write msg error : ", msgID)
		return errors.New("conn write error!")
	}
	return nil
}
