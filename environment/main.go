package main

import (
	. "fmt"
	. "net/http"
	"os"
	"sync"
)

const SECURE_ADDRESS = ":1025"

/**
	Here we've defined a global variable address
	which we set in init() to either the value
	provided in SERVE_HTTP or a default value
	":1024"
**/
var address string
var servers sync.WaitGroup

func init() {
	/**
		We're going to query the evironment for the variable
		SERVE_HTTP which we'll assume contains the default
		address on which to serve unencrypted web content.
	**/
	if address = os.Getenv("SERVE_HTTP"); address == "" {
		address = ":1024"
	}
}

func main() {
	message := "Hello World"

	HandleFunc("/hello", func(w ResponseWriter, r *Request) {
		w.Header().Set("Content-Type", "text/plain")
		Fprintf(w, message)
	})

	Launch(func() {
		ListenAndServe(address, nil)
	})

	Launch(func() {
		ListenAndServeTLS(SECURE_ADDRESS, "cert.pem", "key.pem", nil)
	})

	servers.Wait()
}

func Launch(f func()) {
	servers.Add(1)

	go func() {
		defer servers.Done()
		f()
	}()
}
