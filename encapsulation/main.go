package main

import "fmt"

/*
	Message contains two values: X and y.
	struct describes an area of allocated memory
	which is subdivided into slots for holding
	named values, where each named value has its
	own type.
	For a struct type, every field will be initialised
	to the zero value for their type.
*/
type Message struct {
	X string
	y *string
}

/*
	Print is a method.
	A method is a function that has a defined
	receiver, in OOP terms, a method is a function
	on an instance of an object. The method receiver
	appears in its own argument list between the func
	keyword and the method name.
*/
func (v Message) Print() {
	if v.y != nil {
		fmt.Println(v.X, *v.y)
	} else {
		fmt.Println(v.X)
	}
}

// Store is a method.
func (v *Message) Store(x, y string) {
	v.X = x
	v.y = &y
}

func main() {
	// Literal
	m := &Message{}

	/*
		Called on a Message value to display
		it on terminal.
	*/
	m.Print()

	/*
		Called on a pointer to a Message value
		to change its contents. The reason Store()
		applies to a pointer is that we want to be
		able to change the contents of the Message
		and have these changes persist.
	*/
	m.Store("Hello", "world")
	m.Print()
}
