package main

import (
	. "fmt"

	/*
		net/http package provides a fully-functional
		web server which requires very little configuraiton.
	*/
	"net/http"
)

const MESSAGE = "Hello World"

/*
	ADDRESS is the localhost port and a
	constant that cannot be changed.
	Port 1024 is usually the first non-
	privileged port on most Unix-like
	operating systems.
*/
const ADDRESS = ":1024"

func main() {
	//  http://0.0.0.0:1024/hello
	http.HandleFunc("/hello", Hello)

	/*
		http.ListenAndServe returns an error if it's unable
		to launch the server for some reason, which in this case
		we can print to the console.
	*/
	if e := http.ListenAndServe(ADDRESS, nil); e != nil {
		Println(e)
	}
}

/*
	Hello Handler.
	All we have to do to get our content to the browser is
	define a handler, which in this case is a function to
	call whenever an http.Request is received, and then
	launch a server or listen on the desired address with
	http.ListenAndServe().
*/
func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	Fprintf(w, MESSAGE)
}
