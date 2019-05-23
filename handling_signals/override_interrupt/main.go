package main

import (
	. "fmt"
	. "net/http"
	"os"
	"os/signal"
	. "sync"

	/**
		syscall package defines a number of os.Signal
		values which can be detected by Notify().
	**/
	"syscall"
)

const ADDRESS = ":1024"
const SECURE_ADDRESS = ":1024"

var servers WaitGroup

func init() {
	go SignalHandler(make(chan os.Signal, 1))
}

func main() {
	message := "Hello World"
	HandleFunc("/hello", func(w ResponseWriter, r *Request) {
		w.Header().Set("Content-Type", "text/plain")
		Fprintf(w, message)
	})

	Launch("HTTP", func() error {
		return ListenAndServe(ADDRESS, nil)
	})

	Launch("HTTPS", func() error {
		return ListenAndServeTLS(SECURE_ADDRESS, "cert.pem", "key.pem", nil)
	})

	servers.Wait()
}

/**
	Launch() function takes a name which can
	be displayed as part of an error message,
	and its function parameter now has the signature
	func() error which specifies that it must return
	an error value, which is what's returned by both
	ListenAndServe() and ListenAndServeTLS().

	In the event the error (which is predeclared interface)
	contains a value then we know an error condition's occurred
	and can send a SIGABRT signal with syscall.Kill().
**/
func Launch(name string, f func() error) {
	servers.Add(1)
	go func() {
		defer servers.Done()

		if e := f(); e != nil {
			Println(name, "->", e)
			syscall.Kill(syscall.Getpid(), syscall.SIGABRT)
		}
	}()
}

/**
	SignalHandler() is using a standard for loop statement
	to poll for input from the signal channel and then compare
	it to the cases of a switch statement.

	SIGABRT: Signal Abort.
	SIGTERM: Signal Termination.
	SIGQUIT: Dum core Signal.

	We'll treat SIGABRT as an error condition and the
	SIGTERM and SIGQUIT as clean terminations.

	Sending a SIGABRT signal from our program to itself
	when there's an error launching one of the servers,
	in this case by setting ADDRESS and SECURE_ADDRESS to
	the same value.
**/
func SignalHandler(c chan os.Signal) {
	signal.Notify(c, os.Interrupt, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGQUIT)
	for s := <-c; ; s = <-c {
		switch s {
		case syscall.SIGABRT:
			Println("abnormal exit")
			os.Exit(1)
		case os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT:
			Println("clean shutdown")
			os.Exit(0)
		}
	}
}
