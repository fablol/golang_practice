package main

import (
	"fmt"
	"golang_practice/zinx/znet"
	"io"
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
		// write packed msg
		dp := znet.NewDataPack()
		msg, err := dp.Pack(znet.NewMsgPackage(0, []byte("ZinxV0.5 client test message")))

		if err != nil {
			fmt.Println("pack msg err ", err)
			break
		}

		if _, err := conn.Write(msg); err != nil {
			fmt.Println("write msg err ", err)
			break
		}
		// read head
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, headData); err != nil {
			fmt.Println("read msg head error : ", err)
			break
		}
		// unpack
		msgHead, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack msg head error : ", err)
			break
		}

		if msgHead.GetMsgLen() > 0 {
			// read body
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msg.GetMsgLen())

			if _, err := io.ReadFull(conn, msg.Data); err != nil {
				fmt.Println("read msg body error : ", err)
				break
			}
			fmt.Printf("----> server echo ID: %d, data : %s \n", msg.ID, msg.Data)
		}

		time.Sleep(3 * time.Second)
	}
}
