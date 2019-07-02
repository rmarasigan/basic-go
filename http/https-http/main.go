package main

import (
	. "fmt"
	. "net/http"
)

const ADDRESS = ":1024"
const SECURE_ADDRESS = ":1025"

func main() {
	message := "Hello World!"

	HandleFunc("/hello", func(w ResponseWriter, r *Request) {
		w.Header().Set("Content-Type", "text/plain")
		Fprintf(w, message)
	})

	/*
		Channels are typed conduit through which you can send
		and receive values with the channel operator <-.

		Example:
			ch <- v    (Send v to channel ch)
			v := <-ch  (Receive from ch, and assign value to v)
	*/
	done := make(chan bool)

	/*
		Goroutines run in the same address space, so access to
		shared memory must be synchronized.
	*/
	go func() {
		go func() {
			ListenAndServe(ADDRESS, nil)

			// Send true to channel done.
			done <- true
		}()

		ListenAndServeTLS(SECURE_ADDRESS, "cert.pem", "key.pem", nil)
		done <- true
	}()

	<-done
	<-done
}
