package main

import (
	. "fmt"
	. "net/http"
)

const ADDRESS = ":1024"
const SECURE_ADDRESS = ":1205"

// Launch both of the goroutines from the main() function.
func main() {
	message := "Hello world."

	HandleFunc("/hello", func(w ResponseWriter, r *Request) {
		w.Header().Set("Content-Type", "text/plain")
		Fprint(w, message)
	})

	/*
	 Channel is allowed to transport. No other type
	 is allowed to be transported using the channel.
	 This is a short hand declaration that is also a
	 valid and concise way to define a channel.
	*/
	done := make(chan bool)

	go func() {
		ListenAndServe(ADDRESS, nil)
		/*
		 Write to channel done.
		 The direction of the arrow with respect to the
		 channel specifies whether the data is sent
		 or received.
		*/
		done <- true
	}()

	ListenAndServeTLS(SECURE_ADDRESS, "cert.pem", "key.pem", nil)
	/*
	 We are receiving data from the done channel. This
	 line of code is blocking which means that until
	 some Goroutine writes data to the done channel,
	 the control will not move to the next line of code.

	 Receives data from the done channel but does not
	 use or store that data in any variable.
	*/
	<-done
}
