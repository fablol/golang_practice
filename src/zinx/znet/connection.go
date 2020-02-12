package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

type Connection struct {
	Conn        *net.TCPConn
	ConnID      uint32
	isClosed    bool
	handleAPI   ziface.HandleFunc
	ExitChannel chan bool
}

// init
func NewConnection(conn *net.TCPConn, ConnID uint32, callback_api ziface.HandleFunc) *Connection {
	c := &Connection{
		Conn:        conn,
		ConnID:      ConnID,
		isClosed:    false,
		handleAPI:   callback_api,
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
		buf := make([]byte, 512)
		count, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("read buf err : ", err)
			continue
		}

		err = c.handleAPI(c.Conn, buf, count)

		if err != nil {
			fmt.Println("connection : ", c.ConnID, " handle err : ", err)
		}
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
func (c *Connection) Send(data []byte) error {
	return nil
}
