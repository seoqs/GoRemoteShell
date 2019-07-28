Simple remote access to the command line.
Works at Linux and Windows systems.
It is tested by ParrotOS\Windows 7 |8| 10.

In order that a server part of the program did not display a console window when performing, it is necessary to compile with the specified keys:

go build -ldflags "-s -H windowsgui" windowsServer.go

Also in a server part the encoding/charmap package is used, before compilation it is necessary to establish:

go get "golang.org/x/text/encoding/charmap"
