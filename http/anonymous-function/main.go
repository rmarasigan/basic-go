package main

import (
	. "fmt"
	. "net/http"
)

const MESSAGE = "Hello World!"
const ADDRESS = ":1024"

func main() {
	/*
		HandleFunc() is passed an anonymous function value which
		is created at runtime. This value is a closure so it can
		also access variable in the lexical scope in which it's
		defined.
	*/
	HandleFunc("/hello", func(w ResponseWriter, r *Request) {
		/*
			Header() represents the key-value pairs in an HTTP header.
			Set() sets the header entries associated with key to the
			single element value. It replaces any existing values associated
			with key.
		*/
		w.Header().Set("Content-Type", "text/plain")
		/*
			Fprintf() formats according to a format specifier and writes
			to w. It returns the number of bytes written and any write
			error encountered.
		*/
		Fprintf(w, MESSAGE)
	})

	ListenAndServe(ADDRESS, nil)
}
