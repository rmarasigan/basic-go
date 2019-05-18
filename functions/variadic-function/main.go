package main

import . "fmt"

func main() {
	print("Hello", "world")
}

/**
	interface{} acts as a proxy for any other type
	in a Go program.
	v ...interface{} declare a parameter v which
	takes any number of values.
**/
func print(v ...interface{}) {
	/**
		Println() function is itself defined as
		Println(...interface{}).
	**/
	Println(v...)
}
