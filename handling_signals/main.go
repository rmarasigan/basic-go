package main

import (
	. "fmt"
	. "net/http"
	"os"

	/*
		os/signal package:
		Allows us to register a channel (an atomic queue for
		transferring messages between goroutines at runtime) on
		which notifications are to be received using the
		signal.Notify() function.
	*/
	"os/signal"
	. "sync"
)

const ADDRESS = ":1024"
const SECURE_ADDRESS = ":1025"

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

	Launch(func() {
		ListenAndServe(ADDRESS, nil)
	})

	Launch(func() {
		ListenAndServeTLS(SECURE_ADDRESS, "cert.pem", "key.pem", nil)
	})

	servers.Wait()
}

func Launch(f func()) {
	servers.Add(1)
	go func() {
		defer servers.Done()
		f()
	}()
}

/*
	SignalHandler() consists of a goroutine
	containing an infinite loop and blocking on
	a channel of fixed size (in this case able to
	hold only one element at a time).

	SignalHandler() should be trapping the Interrupt
	and Kill signals. In both cases if we catch the
	signal we print a message to the console before
	exiting, however as previously mentioned the Kill
	signal (which can be sent from another shell session
	using the kill command) will never be received by our
	Go code.
*/
func SignalHandler(c chan os.Signal) {
	signal.Notify(c, os.Interrupt)

	for s := <-c; ; s = <-c {
		switch s {
		/*
			os.Interrupt can be sent with
			control-C.
		*/
		case os.Interrupt:
			Println("^C received")
			os.Exit(0)

			/*
				os.Kill equates to SIGKILL on
				*nixen and is usually a non-maskable
				interrupt, meaning that it terminates
				execution and will never be received
				by Notify().

				Kill signal (which can be sent from another
				shell session using the kill command) will
				never be received by our Go code.
			*/
		case os.Kill:
			Println("SIGKILL received")
			os.Exit(1)
		}
	}
}
