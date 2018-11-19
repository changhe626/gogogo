package first

import (
	"fmt"
	"net"
)

func main() {

	fmt.Println("start server")

	listen, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("error listening", err.Error())
		return //终止程序
	}
	//监听来自客户端的链接
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("error accept", err.Error())
			return
		}
		go deServerStuff(conn)
	}

}
func deServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("error readinf", err.Error())
			return
		}
		fmt.Println("received data:%v", string(buf))
	}

}
