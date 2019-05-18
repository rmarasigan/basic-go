package main

import . "fmt"

const Hello = "Hello"

func main() {
	Println(Hello, world())
}

/**
	The empty brackets () indicate that there are
	no parameters passed into the function when it's called.
**/
func world() string {
	/**
		return is required by the language specification
		whenever a function specifies return values.
	**/
	return "world"
}

// The () first for parameters and second for results.
func worldLong() string {
	return "world"
}
