package main

/*
	How to run our program:
		$ go run main.go &
		[1] 29137

		Note: telnet allows us to connect directly to
		a TCP/IP server and interact with it.

		$ telnet localhost 1024
		Trying 127.0.0.1...
		Connected to localhost.
		Escape character is '^]'.
		Hello World
		Connection closed by foreign host.
*/

import (
	. "fmt"
	/*
		net package revolves around server the Listener
		and client Connection types.

		Listener: is an interface which allows any type
		implementing its specified methods - Accept(),
		Close() and Addr() - to be used interchangeably
		and is a key tool in Go for generalising program
		design.

		* Listen() on a specified protocol and port number.
		* Accept() incoming connections.
		* For each net.Conn, run a handler in a separate goroutine.
		* Read from and write to the connection while performing work.
		* Close each connection when it finishes its work.
	*/
	"net"
)

func main() {
	if listener, e := net.Listen("tcp", ":1024"); e == nil {
		for {
			if connection, e := listener.Accept(); e == nil {
				/*
					net.Conn implements the Writer interface defined in the io
					package.

					fmt.Fprintf takes any type which integrates io.Writer as its
					target.
				*/
				go func(c net.Conn) {
					defer c.Close()
					Fprintln(c, "Hello World")
				}(connection)
			}
		}
	}
}
