package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	fmt.Print("Client for simple remote shell...\n")
	readIP := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Server IP address: ")

	strIPAddress, _ := readIP.ReadString('\n')
	strIPAddress = strings.TrimSpace(strIPAddress)

	conn, err := net.Dial("tcp", strIPAddress+":8087")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error dialing tcp: %v\nServer needs to be runing before client.\n", err)
		return
	}
	defer conn.Close()
	console := bufio.NewReader(os.Stdin)
	connReader := bufio.NewReader(conn)

	for {
		fmt.Print("For close connection clien use command: client close\n")
		fmt.Print("For stop server use command: server close\n")
		fmt.Print("Command to send: ")
		text, err := console.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading string: %v", err)
			return
		}

		text = strings.TrimSpace(text)

		fmt.Fprintf(conn, text+"\n")
		if text == "client close" {
			fmt.Println("Close connection")
			return
		}
		if text == "server close" {
			fmt.Println("Close connection")
			return
		}
		for {
			message, err := connReader.ReadString('\n')
			if err != nil {
				fmt.Fprintf(os.Stderr, "error reading from conn: %v", err)
				return
			}

			if strings.Index(message, "<<<endMessage>>") > 0 {
				break
			}
			fmt.Print(message)
		}
	}

}
