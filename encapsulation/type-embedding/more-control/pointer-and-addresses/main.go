package main

import . "fmt"

type Text string

func main() {
	/**
		Go has a means of referring to specific
		addresses in memory and of accessing their
		contents directly.
		Go allows user-defined types to declare
		methods on either a value type or a
		pointer to a value type.
	**/

	var name Text = "Ellie"
	var pointerToName *Text

	pointerToName = &name

	Printf("name = %v stored at %v\n", name, pointerToName)
	Printf("pointerToName references %v\n", *pointerToName)
}
