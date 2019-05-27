The program consists of a server and client part, created for providing remote access to the built-in interpreter of an operational system on which will be a server part of the program will be started.

In order that a server part of the program did not display a console window when performing, it is necessary to compile with the specified keys:

go build -ldflags "-s -H windowsgui" winServerShell.go

This version of the program can be carried out both on the operational Windows systems, and on * by Unix similar systems.

In case of use of a server part on * Unix similar systems needs to replace an interpreter name in the 37th line:

"with cmd: = exec.Command ("cmd", argsCommand...)" 

"cmd" on "bash | sh".

And in a client part to correct one symbol in the 19th line:

"fmt.Fprintf (conn, " / c " text " \n")"

"/ c " on "- c ".


Also in a server part the encoding/charmap package is used, before compilation it is necessary to establish:

go get "golang.org/x/text/encoding/charmap"
