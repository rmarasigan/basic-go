/*
	package is a library which provides a group of related
	functions and data types.
*/
package main

/*
	import statement is a reference to the fmt package.
	import . "fmt" allows the explicit package reference to be
	removed from the Println() function call.
*/
import (
	/*
		fmt provides functions and types associated with
		formatting text for printing and displaying it on
		a console or in the command shell.
	*/
	"fmt"
)

// main() function neither takes parameters nor has a return value
func main() {
	// println() - usually used to assist debugging
	println("Hello World!")

	/*
		fmt.Println() takes one or more parameters and prints
		them to the console with a carriage return appended.
	*/
	fmt.Println("HELLO WORLD!")
}
