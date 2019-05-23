package main

import (
	. "fmt"
	. "net/http"
	"os"
	"sync"
)

var (
	address       string
	secureAddress string
	certificate   string
	key           string
)

var servers sync.WaitGroup

func init() {

	// To make the program fully configurable from the environment.
	if address = os.Getenv("SERVE_HTTP"); address == "" {
		address = ":1024"
	}

	if secureAddress = os.Getenv("SERVE_HTTPS"); secureAddress == "" {
		secureAddress = ":1025"
	}

	if certificate = os.Getenv("SERVE_CERT"); certificate == "" {
		certificate = "cert.pem"
	}

	if key = os.Getenv("SERVE_KEY"); key == "" {
		key = "key.pem"
	}
}

func main() {
	message := "Hello once again, world!"

	HandleFunc("/hello", func(w ResponseWriter, r *Request) {
		w.Header().Set("Content-Type", "text/plain")
		Fprint(w, message)
	})

	Launch(func() {
		ListenAndServe(address, nil)
	})

	Launch(func() {
		ListenAndServeTLS(secureAddress, certificate, key, nil)
	})

	servers.Wait()
}

func Launch(f func()) {
	servers.Add(1)

	go func() {
		defer servers.Done()
	}()
}
