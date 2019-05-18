package main

import (
	. "fmt"
)

func main() {
	// Prints hello world.
	greet("world")
}

// greet method returns one value which is name.
func greet(name string) {
	Println("hello", name)
}
