package main

import . "fmt"

const Hello = "Hello"

var world string

/*
	When there is a multiple init() functions
	the order in which they're executed is
	indeterminate so in general it's best not
	to do this unless you can be certain the
	init() functions don't interact in any way.
*/
func init() {
	Print(Hello, " ")
	world = "world"
}

func init() {
	Printf("%v\n", world)
}

func main() {}
