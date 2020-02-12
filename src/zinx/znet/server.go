package znet

import (
	"errors"
	"fmt"
	"net"
	"zinx/ziface"
)

/***
 *                    .::::.
 *                  .::::::::.
 *                 :::::::::::  FUCK YOU
 *             ..:::::::::::'
 *           '::::::::::::'
 *             .::::::::::
 *        '::::::::::::::..
 *             ..::::::::::::.
 *           ``::::::::::::::::
 *            ::::``:::::::::'        .:::.
 *           ::::'   ':::::'       .::::::::.
 *         .::::'      ::::     .:::::::'::::.
 *        .:::'       :::::  .:::::::::' ':::::.
 *       .::'        :::::.:::::::::'      ':::::.
 *      .::'         ::::::::::::::'         ``::::.
 *  ...:::           ::::::::::::'              ``::.
 * ```` ':.          ':::::::::'                  ::::..
 *                    '.:::::'                    ':'````..
 */

/***
 * ░░░░░░░░░░░░░░░░░░░░░░░░▄░░
 * ░░░░░░░░░▐█░░░░░░░░░░░▄▀▒▌░
 * ░░░░░░░░▐▀▒█░░░░░░░░▄▀▒▒▒▐
 * ░░░░░░░▐▄▀▒▒▀▀▀▀▄▄▄▀▒▒▒▒▒▐
 * ░░░░░▄▄▀▒░▒▒▒▒▒▒▒▒▒█▒▒▄█▒▐
 * ░░░▄▀▒▒▒░░░▒▒▒░░░▒▒▒▀██▀▒▌
 * ░░▐▒▒▒▄▄▒▒▒▒░░░▒▒▒▒▒▒▒▀▄▒▒
 * ░░▌░░▌█▀▒▒▒▒▒▄▀█▄▒▒▒▒▒▒▒█▒▐
 * ░▐░░░▒▒▒▒▒▒▒▒▌██▀▒▒░░░▒▒▒▀▄
 * ░▌░▒▄██▄▒▒▒▒▒▒▒▒▒░░░░░░▒▒▒▒
 * ▀▒▀▐▄█▄█▌▄░▀▒▒░░░░░░░░░░▒▒▒
 * 单身狗就这样默默地看着你，一句话也不说。
 */

// server moudel
type Server struct {
	// name
	Name string
	// ip version
	IPVersion string
	// ip
	IP string
	// port
	Port int
}

func EchoToClient(conn *net.TCPConn, data []byte, count int) error {
	fmt.Printf("[EchoToClient] EchoToClient.... \n")
	if _, err := conn.Write(data[:count]); err != nil {
		fmt.Println("[EchoToClient]write data err", err)
		return errors.New("EchoToClient")
	}
	return nil
}

func (s *Server) Start() {
	// get tcp addr
	fmt.Printf("[Start] Server Listerner at IP :%s, Port %d, is starting\n", s.IP, s.Port)

	go func() {
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error : ", err)
			return
		}
		// listen
		listerner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen ", s.IPVersion, " err ", err)
			return
		}
		fmt.Println("start Zinx server sucess, ", s.Name, "listenning")
		// wait connect,handle

		var cid uint32
		cid = 0
		for {
			conn, err := listerner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}
			// do
			// echo 512byte message
			// go func() {
			// 	for {
			// 		buf := make([]byte, 512)
			// 		count, err := conn.Read(buf)
			// 		if err != nil {
			// 			fmt.Println("recv buf err", err)
			// 		}

			// 		//echo
			// 		fmt.Printf("server recv : %s, count : %d \n", buf, count)
			// 		if _, err := conn.Write(buf[:count]); err != nil {
			// 			fmt.Println("echo buff err", err)
			// 			continue
			// 		}
			// 	}
			// }()

			handle_conn := NewConnection(conn, cid, EchoToClient)
			cid++
			// start
			go handle_conn.Start()
		}
	}()
}

func (s *Server) Stop() {
	// todo clear
}

func (s *Server) Serve() {
	s.Start()
	//
	select {}
}

// init
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
	return s
}
