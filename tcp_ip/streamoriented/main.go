package main

import (
	"bufio"
	. "fmt"
	"net"
)

func main() {
	/**
		For stream-oriented protocols like TCP/IP we do this using
		the net.Dial() function to open a net.Conn connection to a
		server and we can then interact with this using the io.Reader
		and io.Writer interfaces.

		net.Conn represents streams of data flowing between client and
		server we've introduced the bufio package to our client so that
		the data it's receiving is buffered.
	**/
	if connection, e := net.Dial("tcp", ":1024"); e == nil {
		defer connection.Close()

		if text, e := bufio.NewReader(connection).ReadString('\n'); e == nil {
			Printf(text)
		}
	}
}
