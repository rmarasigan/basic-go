package main

import (
	. "fmt"
	. "net/http"
)

const SECURE_ADDRESS = ":1205"

/*
	Before we run the program, we first need to generate a
	certificate and a public key. This is a self-signed
	certificate, and not all modern web browsers like these
	For production applications you'll need a certificate from
	a recognised Certificate Authority.
*/

func main() {
	message := "Hello world"

	HandleFunc("/hello", func(w ResponseWriter, r *Request) {
		w.Header().Set("Content-Type", "text/plain")
		Fprintf(w, message)
	})

	/*
		ListenAndServeTLS acts identically to ListenAndServe, except
		that it expects HTTPS connections. Additionally, files containing
		a certificate and matching private key for the server must be
		provided.
	*/
	ListenAndServeTLS(SECURE_ADDRESS, "cert.pem", "key.pem", nil)
}
