package main

import "fmt"

/*
	In this case HelloWorld{} is just an empty struct, but
	which in reality could be any declared type.
*/
type HelloWorld struct{}

/*
	HelloWorld defines a String() method which
	can be called on any HelloWorld value.
*/
func (h HelloWorld) String() string {
	return "Hello world"
}

/*
	Message which embeds the HelloWorld type by
	defining an anonymous field of the HelloWorld type.
*/
type Message struct {
	HelloWorld
}

func main() {
	m := &Message{}
	fmt.Println(m.HelloWorld.String())
	fmt.Println(m.String())
	fmt.Println(m)
}
