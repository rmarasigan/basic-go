package main

import . "fmt"

const Hello = "Hello"

var world string

/*
	Every Go package may contain one or more init()
	functions specifying actions that should be taken
	during program initialisation.

	We use the init() function to assign to our world variable.
*/
func init() {
	world = "world"
}

func main() {
	Println(Hello, world)
}
