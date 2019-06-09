package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	readIP := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Server IP address: ")
	strIPAddress, _ := readIP.ReadString('\n')
	servIPAddress := strings.TrimRight(strIPAddress, "\r\n")
	conn, _ := net.Dial("tcp", servIPAddress+":8081")
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Command to send: ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")
		connbuf := bufio.NewReader(conn)
		for {
			str, _ := connbuf.ReadString('\n')
			if strings.Index(str, "<<<endMessage>>") > 0 {
				break
			}
			fmt.Print(str)
		}
	}
	conn.Close()
}
