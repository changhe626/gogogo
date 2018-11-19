package first

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	//打开链接
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("error diaaling", err.Error())
		return
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Firsr, what is your name?")
	clientName, _ := reader.ReadString('\n')
	trimClient := strings.Trim(clientName, "\r\n")
	//给服务器发送消息知道程序退出
	for {
		fmt.Println("what to send to server?Type Q to quit.")
		input, _ := reader.ReadString('\n')
		trimIput := strings.Trim(input, "\r\n")
		if trimIput == "Q" {
			return
		}
		conn.Write([]byte(trimClient + "says" + trimIput))

	}

}
