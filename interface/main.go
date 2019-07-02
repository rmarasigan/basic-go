package main

import "fmt"

/*
	interface{} is a collection of method signatures
	that an Object can implement.

	Like struct, we need to use type alias to simplify
	interface declaration along with keyword interface.
*/
type Stringer interface {
	/*
		Stringer interface which has 1 method, String()
		which accepts no arguments and return string.
	*/
	String() string
}

// Empty struct will occupy zero bytes of storage.
type Hello struct{}

// Method String of Hello.
func (h Hello) String() string {
	return "Hello"
}

// Empty World struct.
type World struct{}

// Method String of World.
func (w *World) String() string {
	return "world"
}

/*
	struct are typed collections of fields.
	Message struct having X and Y fields and
	has a type of string.
*/
type Message struct {
	X Stringer
	Y Stringer
}

func (v Message) String() (r string) {
	switch {
	case v.X == nil && v.Y == nil:
	case v.X == nil:
		r = v.Y.String()
	case v.Y == nil:
		r = v.X.String()
	default:
		r = fmt.Sprintf("%v %v", v.X, v.Y)
	}

	return
}

func main() {
	m := &Message{}
	fmt.Println(m)

	m.X = new(Hello)
	fmt.Println(m)

	m.Y = new(World)
	fmt.Println(m)

	m.Y = m.X
	fmt.Println(m)

	m = &Message{X: new(World), Y: new(Hello)}
	fmt.Println(m)

	m.X, m.Y = m.Y, m.X
	fmt.Println(m)
}
