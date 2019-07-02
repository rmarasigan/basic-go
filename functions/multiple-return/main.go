package main

import (
	. "fmt"
)

func main() {
	/*
		Println() happens to be a variadic function,
		and takes zero or more parameters so it happily,
		consumes both of the values message() returns.
	*/
	Println(message())
}

/*
	message() returns two values, we can use it
	in any context where at least two parameters
	can be consumed.
*/
func message() (string, string) {
	return "hello", "world"
}
