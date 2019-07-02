package main

import (
	. "fmt"
	. "net/http"
)

const MESSAGE = "Hello World"
const ADDRESS = ":1024"

func main() {
	/*
		HandleFunc() registers a URL in the web server
		as the trigger for a function, so when a web request
		targets the URL the associated function will be executed
		to generate the result.
	*/
	HandleFunc("/hello", Hello)

	/*
		ListenAndServe() which will block for as long as the server
		is active, returning an error if there is one to report.
	*/
	ListenAndServe(ADDRESS, nil)
}

/*
	Hello Handler
	ResponseWriter to send output to the web client.
	The ResponseWriter is a file handle so we can use
	fmt.Fprint() family.
	Request is being replied to.
*/
func Hello(w ResponseWriter, r *Request) {
	w.Header().Set("Content-Type", "text/plain")

	/*
		fmt.Fprintf() - file-writing functions to
		create the response body.
		fmt.Fprintf() formats according to a format
		specifier and writes to w.
	*/
	Fprintf(w, MESSAGE)
}
