package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("client start...")
	// connect
	time.Sleep(1 * time.Second)
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("clent start err ", err)
		return
	}

	for {
		// write
		_, err := conn.Write([]byte("hello, sb"))
		if err != nil {
			fmt.Println("write buff err ", err)
			return
		}

		buf := make([]byte, 512)
		count, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buff err ", err)
			return
		}
		fmt.Printf("server echo : %s, count : %d \n", buf, count)

		time.Sleep(3 * time.Second)
	}
}
