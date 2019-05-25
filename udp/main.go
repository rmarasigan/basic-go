package main

/**
	UDP is an unreliable transport dealing in
	individual packets (datagrams) which are
	independent of each other, therefore a server
	doesn't maintain streams to any of its clients
	and these are responsible for any error-correction
	or packet ordering which may be necessary to
	coordinate successive signals.
**/

import (
	. "fmt"
	"net"
)

var HELLO_WORLD = ([]byte)("Hello World\n")

func main() {
	/**
		net.ResolveUDPAddr() - to resolve the address.
		net.ListenUDP() - opens a UDP port and listens
		for traffic.
		net.ReadFromUDP() - copies incoming data into a
		buffer and provides the remote client's address.
		net.WriteToUDP() - writes data back to the remote
		client's address.
	**/
	if address, e := net.ResolveUDPAddr("udp", ":1024"); e == nil {
		if server, e := net.ListenUDP("udp", address); e == nil {
			for buffer := MakeBuffer(); ; buffer = MakeBuffer() {
				if n, client, e := server.ReadFromUDP(buffer); e == nil {
					go func(c *net.UDPAddr, packet []byte) {
						if n, e := server.WriteToUDP(HELLO_WORLD, c); e == nil {
							Printf("%v bytes written to: %v\n", n, c)
						}
					}(client, buffer[:n])
				}
			}
		}
	}
}

func MakeBuffer() (r []byte) {
	return make([]byte, 1024)
}
