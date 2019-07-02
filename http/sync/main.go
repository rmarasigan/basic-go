package main

import (
	. "fmt"
	. "net/http"
	"sync"
)

const ADDRESS = ":1024"
const SECURE_ADDRESS = ":1025"

func main() {
	message := "Hello World! :)"

	HandleFunc("/hello", func(w ResponseWriter, r *Request) {
		w.Header().Set("Content-Type", "text/plain")
		Fprintln(w, message)
	})

	/*
		WaitGroup waits for a collection of goroutines to finish.

		sync.WaitGroup to keep track of how many active
		goroutines we have in our program and only terminate
		the program when they've all completed their work.
		Again, this allows us to run both servers in separate
		goroutines and manage termination correctly.
	*/
	var servers sync.WaitGroup

	/*
		Main goroutine calls Add to set the number of goroutines
		to wait for. Then each of the goroutines runs and calls
		Done when finished.
	*/
	servers.Add(1)

	go func() {
		/*
			Indicate that goroutine is finished executing to
			the WaitGroup.
		*/
		defer servers.Done()
		ListenAndServe(ADDRESS, nil)
	}()

	servers.Add(1)

	go func() {
		defer servers.Done()
		ListenAndServeTLS(SECURE_ADDRESS, "cert.pem", "key.pem", nil)
	}()

	/*
		Wait can be used to block until all goroutines
		have finished.
	*/
	servers.Wait()
}
