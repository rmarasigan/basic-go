package main

import (
	. "fmt"
	. "net/http"
)

const ADDRESS = ":1024"

/**
	A basic premise by defining a variable messages in
	main() and then accessing it from within the
	anonymous function.
**/
func main() {
	message := "Hello World"
	HandleFunc("/hello", func(w ResponseWriter, r *Request) {
		w.Header().Set("Content-Type", "text/plain")
		Fprint(w, message)
	})

	ListenAndServe(ADDRESS, nil)
}
