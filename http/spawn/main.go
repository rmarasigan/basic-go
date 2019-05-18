package main

import (
	. "fmt"
	. "net/http"
)

const ADDRESS = ":1024"
const SECURE_ADDRESS = ":1025"

func main() {
	message := "Hello once again, World."
	HandleFunc("/hello", func(w ResponseWriter, r *Request) {
		w.Header().Set("Content-Type", "text/plain")
		Fprint(w, message)
	})

	Spawn(
		func() {
			ListenAndServeTLS(SECURE_ADDRESS, "cert.pem", "key.pem", nil)
		},

		func() {
			ListenAndServe(ADDRESS, nil)
		},
	)
}

/**
	Spawn.
	Allowing us to run any arbitrary piece of code and
	wait for it to signal completion. It's also a variadic
	function, taking as many or as few functions as desired
	and setting each of them up correctly.
**/
func Spawn(f ...func()) {
	done := make(chan bool)

	for _, s := range f {
		/**
			By accepting the parameter server to the goroutine's
			closure we can pass in the value of s and capture it
			so that on successive iterations of the range our
			goroutines use the correct value.
		**/
		go func(server func()) {
			server()
			done <- true
		}(s)
	}

	for l := len(f); l > 0; l-- {
		<-done
	}
}
