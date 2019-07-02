package main

import . "fmt"

/*
	const - value cannot be changed at runtime.
	The identifier Hello starts with a capital
	letter the associated constant is visible to
	other packages - though this isn't releveant
	in the context of a main package.
	The identifier world starts with a lowercase
	letter and is only accessible within the main
	package.
*/
const Hello = "Hello"
const world = "world"

/*
	Declaring a variable and assigning a value to it.
	If we wish to be more explicit we can use:
	var variableWorld string = "world"
	Global Variable: variableWorld
*/
var variableWorld = "world"

func main() {

	/*
		If we chose to declare the local variable separately
		from the assignment we'd have to give it a different
		name to avoid a compilation error.
		When w is declared it's also initialised to the zero
		value, which in the case of string happens to be "".
	*/
	var w string
	w = variableWorld + "!"

	/*
		New local variable which is variableWorld within
		main() which takes its value from an operation
		concatenating the value of the global variableWorld
		variable with an exclamation mark.
		Shadowing - within main(), any subsequent reference
		to variableWorld will always access the local version
		of  the variable without affecting the global variableWorld
		variable.
	*/
	variableWorld += "!"

	/*
		Println() is variadic we can pass it both
		constants and it will print them on the same
		line, separate by whitespace.
	*/
	Println(Hello, world)

	/*
		Printf() gives precise control over how its
		parameters are displayed using a format
		specifier.
		% - determine how a particular parameter will be
		displayed.
		%v - it allows the formatting to be specified by
		the type of the parameter.
		\n - mark a new line.
	*/
	Printf("%v %v \n", Hello, world)

	Println(Hello, variableWorld)

	Println(Hello, w)

}
