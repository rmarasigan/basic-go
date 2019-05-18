package main

import "fmt"

type Hello struct{}

func (h Hello) String() string {
	return "Hello"
}

type Message struct {
	/**
		Message type embeds a Hello by-pointer.
		The first advantage to this is that you rely on functions
		that use the NewX idiom returning a struct by-pointer to do
		initialization.
		The second advantage is that you can embed all the functionality
		of a type without needing to know when it is instantiated.
	**/
	*Hello
	World string
}

func (v Message) String() (r string) {
	if v.Hello == nil {
		r = v.World
	} else {
		/**
			fmt.Sprintf() formats according to a format specifier and
			returns the resulting string.
		**/
		r = fmt.Sprintf("%v %v", v.Hello, v.World)
	}

	return
}

func main() {
	/**
		& operator is prepended to the name of a variable
		or to a value literal when we wish to discover its
		address in memory, which we refer to as a pointer.
		To do anything with the pointer returned by the &
		operator we need to be able to declare a pointer
		variable which we do by prepending a type name with
		the * operator.
	**/
	m := &Message{}
	fmt.Println(m)

	m.Hello = new(Hello)
	fmt.Println(m)

	m.World = "world"
	fmt.Println(m)
}
