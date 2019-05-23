package main

import (

	/**
		For a server we import "crypto/rand" to access
		rand.Reader, a cryptographically secure pseudo-
		random number generator which we'll be using as
		a source of randomness in the TLS connection.
	**/
	"crypto/rand"

	/**
		"crypto/tls" provides us with an equivalent API
		to that defined in net and this means that as
		tls.Listen() fulfils the net.Listener interface
		our connections will be of type net.Conn.
	**/
	"crypto/tls"
	. "fmt"
)

func main() {
	/**
		We then create a certificate using tls.LoadX509KeyPair()
		to load the server key pair and if this is successful then
		we set up a listener to accept incoming connections and
		write "Hello everyone!" to a client.
	**/
	if certificate, e := tls.LoadX509KeyPair("server.cert.pem", "server.key.pem"); e == nil {
		config := tls.Config{
			Certificates: []tls.Certificate{certificate},
			Rand:         rand.Reader,
		}

		if listener, e := tls.Listen("tcp", ":1205", &config); e == nil {
			for {
				if connection, e := listener.Accept(); e == nil {
					go func(c *tls.Conn) {
						defer c.Close()
						Fprintln(c, "Hello everyone!")
					}(connection.(*tls.Conn))
				}
			}
		}
	}
}
