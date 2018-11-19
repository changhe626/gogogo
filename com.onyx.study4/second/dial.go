package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	// tcp ipv4
	conn, e := net.Dial("tcp", "127.0.0.1:8000")
	checkConnect(conn, e)
	// udp
	conn2, e := net.Dial("udp", "127.0.0.1:8001")
	checkConnect(conn2, e)
	// tcp ipv6
	conn, e = net.Dial("tcp", "[2620:0:2d0:200::10]:80")
	checkConnect(conn, e)

}
func checkConnect(conn net.Conn, e error) {
	if e != nil {
		fmt.Printf("error %v connecting!")
		os.Exit(1)
	}
	fmt.Println("Connection is made with %v", conn)

}
