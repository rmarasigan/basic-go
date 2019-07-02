package main

import (
	"bufio"
	. "fmt"
	"net"
)

var CRLF = ([]byte)("\n")

/*
	Note how each time we run the client program
	it's assigned a different network port by the
	operating system each time net.DialUDP is called.
	In the case of OSX this port will usually be at the
	upper end of the non-privileged port range.
*/

func main() {
	/*
		net.ResolveUDPAddr(network, address string)
		Returns na address of UDP end point. The network
		must be a UDP network name.

		net.ListenUDP(network string, laddr *UDPAddr)
		Acts like ListenPacket for UDP networks. The
		network must be a UDP network name.
	*/
	if address, e := net.ResolveUDPAddr("udp", ":1024"); e == nil {
		if server, e := net.DialUDP("udp", nil, address); e == nil {
			defer server.Close()
			/*
				We can make this apparent by performing
				multiple sequences of Write()
			*/
			for i := 0; i < 3; i++ {
				if _, e = server.Write(CRLF); e == nil {
					if text, e := bufio.NewReader(server).ReadString('\n'); e == nil {
						Printf("%v", text)
					}
				}
			}
		}
	}
}
