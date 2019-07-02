package main

import . "fmt"

const Hello = "hello"

var world string

/*
	Allowing us to place the whole of our
	program in init()
*/
func init() {
	world = "world"
	Println(Hello, world)
}

/*
	Leaving main() as a stub to convince that
	this is indeed a valid Go program.
*/
func main() {}
