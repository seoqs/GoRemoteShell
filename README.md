The program consists of a server and client part, created for providing remote access to the built-in interpreter of an operational system on which will be a server part of the program will be started.

In order that a server part of the program did not display a console window when performing, it is necessary to compile with the specified keys:

go build -ldflags "-s -H windowsgui" winServerShell.go

Also in a server part the encoding/charmap package is used, before compilation it is necessary to establish:

go get "golang.org/x/text/encoding/charmap"
