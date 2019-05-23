package main

import (
	"crypto/rand"
	"crypto/tls"
	. "fmt"
	"net"
	"sync"
)

var servers sync.WaitGroup

func main() {
	if listener, e := net.Listen("tcp", ":1024"); e == nil {
		Serve(listener)
	}

	Serve(TLSListener("server.cert.pem", "server.key.pem", ":1025"))
	servers.Wait()
}

/**
	TLSListener() returns a net.Listener value as
	tls.Listener complies with its interface, or a nil
	value if tls.Listen() returns an error.
**/
func TLSListener(cert, key, address string) (r net.Listener) {
	if certificate, e := tls.LoadX509KeyPair(cert, key); e == nil {
		config := tls.Config{
			Certificates: []tls.Certificate{certificate},
			Rand:         rand.Reader,
		}

		if listener, e := tls.Listen("tcp", address, &config); e == nil {
			r = listener
		}
	}
	return
}

/**
	Serve() - to phrase the server behaviour in terms
	of the net.Listener interface.
**/
func Serve(listener net.Listener) {
	if listener != nil {
		Launch(func() {
			for {
				if connection, e := listener.Accept(); e == nil {
					go func(c net.Conn) {
						defer c.Close()
						Fprintln(c, "Hello World")
					}(connection)
				}
			}
		})
	}
}

/**
	Launch() - to manage the lifecycle of our two server
	goroutines.
**/
func Launch(f func()) {
	servers.Add(1)
	go func() {
		defer servers.Done()
		f()
	}()
}
