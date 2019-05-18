package main

import "fmt"

/**
	We've been using pointers to Hello and World
	so the interface variables are storing pointers
	to pointers to these values (i.e. **Hello and
	**World) rather than pointers to the values
	themselves (i.e. *Hello and *World).

	interfaces support embedding of other interfaces.
**/

type Hello struct{}

func (h Hello) String() string {
	return "Hello"
}

type World struct{}

func (w *World) String() string {
	return "world"
}

type Message struct {
	/**
		fmt.Stringer is implemented by any value that has a String method,
		which defines the “native” format for that value (String() string).
		It is used to print values passed as an operand to any format that
		accepts a string or to an unformatted printer such as Print.
	**/
	X fmt.Stringer
	Y fmt.Stringer
}

/**
	IsGreeting() is a predicate which uses a pair of type
	assertions to tell us whether or not one of Message's
	data fields contains a value of concrete type Hello.
**/
func (v Message) IsGreeting() (ok bool) {
	if _, ok = v.X.(*Hello); !ok {
		_, ok = v.Y.(*Hello)
	}
	return
}

func main() {
	m := &Message{}
	fmt.Println(m.IsGreeting())

	m.X = new(Hello)
	fmt.Println(m.IsGreeting())

	m.Y = new(World)
	fmt.Println(m.IsGreeting())

	m.Y = m.X
	fmt.Println(m.IsGreeting())

	m = &Message{X: new(World), Y: new(Hello)}
	fmt.Println(m.IsGreeting())

	m.X, m.Y = m.Y, m.X
	fmt.Println(m.IsGreeting())
}
