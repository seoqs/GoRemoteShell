package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var servIPAddress = "127.0.0.1"

func main() {
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
