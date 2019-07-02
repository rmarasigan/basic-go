package main

import "fmt"

/*
	Type Aliasing
	Syntax: type aliasName aliasTo

	aliasName is the name you want to give
	it to your type which will behave like
	aliasTo type.
*/
type HelloWorld bool

func (h HelloWorld) String() (r string) {
	/*
		If true, then give r a value that
		corresponds to its type.
	*/
	if h {
		r = "Hello world"
	}

	return
}

type Message struct {
	HelloWorld
}

func main() {

	// Setting HelloWorld boolean value to true.
	m := &Message{HelloWorld: true}
	fmt.Println(m)

	// Shorthand in setting boolean value of HelloWorld.
	m.HelloWorld = false
	fmt.Println(m)

	m.HelloWorld = true
	fmt.Println(m)
}
