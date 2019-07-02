package main

import (
	. "fmt"
	. "net/http"
)

/*
	ADDRESS for HTTP.
	Definitely transfer images and text and even
	sound, but not videos.
*/
const ADDRESS = ":1204"

/*
	SECURE_ADDRESS for HTTPS.
	All communications between your browser and
	the website are encrypted.
*/
const SECURE_ADDRESS = ":1205"

func main() {
	message := "Hello World"

	HandleFunc("/hello", func(w ResponseWriter, r *Request) {
		w.Header().Set("Content-Type", "text/plain")
		Fprintf(w, message)
	})

	/*
		go keyword marks a goroutine which is a lightweight
		thread scheduled by the Go runtime. When a goroutine
		is launched it takes a function call and creates a
		separate thread of execution for it.

		This unfortunately doesn't work either because of the
		way goroutines are scheduled: the infinite loop can
		potentially starve the thread shceduler and prevent the
		other goroutines from running.
	*/
	go func() {
		ListenAndServe(ADDRESS, nil)
	}()

	go func() {
		ListenAndServeTLS(SECURE_ADDRESS, "cert.pem", "key.pem", nil)
	}()

	/*
		We're using an infite for loop to prevent program
		termination: it's inelegant, but this is a small
		program and dirty hacks have their appeal.
	*/
	for {
	}
}
