package main

import (
	. "fmt"
	. "net/http"
	"sync"
)

const ADDRESS = ":1024"
const SECURE_ADDRESS = ":1025"

/**
	Has been turned into a global variable
	to simplify the function of Launch().
**/
var servers sync.WaitGroup

func main() {
	message := "Hello world!!"

	HandleFunc("/hello", func(w ResponseWriter, r *Request) {
		w.Header().Set("Content-Type", "text/plain")
		Fprintf(w, message)
	})

	/**
		When we call Launch() we're freed from the need to
		manually increment servers prior to goroutine startup.
	**/
	Launch(func() {
		ListenAndServe(ADDRESS, nil)
	})

	Launch(func() {
		ListenAndServeTLS(SECURE_ADDRESS, "cert.pem", "key.pem", nil)
	})

	servers.Wait()
}

/**
 Launch takes a parameter-less function and
 wraps this in a closure which will be launched
 as a goroutine in a separate thread of execution.
**/
func Launch(f func()) {
	servers.Add(1)
	go func() {
		/**
		 We use defer statement to automatically call
		 servers.Done() when goroutine terminates
		 even in the event that the goroutine crashes.
		**/
		defer servers.Done()
		f()
	}()
}
