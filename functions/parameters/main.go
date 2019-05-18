package main

import "fmt"

// main() doesn't have a return statement.
func main() {
	fmt.Println(message("world"))
}

/**
	Passing the parameter name which is a string
	value into the function message(), and assigning
	the function's return value to message which is a
	variable declared and available throughout the
	function.

	message() performs a calculation using the
	Sprintf() function and returns its result.
**/
func message(name string) (message string) {
	/**
		Sprintf() is similar to Printf(), only
		rather than create a string according to
		a format and displaying it in the terminal
		instead returns this string as a value which
		can assign to a variable or use as a parameter
		in another function call such as Prntln().
	**/
	message = fmt.Sprintf("hello %v", name)

	/**
		Because we've explicitly named the return value
		we don't need to reference it in the return
		statement as each of the named return values
		is applied.
	**/
	return
}
