package main

import "fmt"

func main() { //entry point
	x := 42

	// Let's print the types of values
	printValueType(x)
	greet := "Hello!"
	printValueType(greet)
	f := 3.14159
	printValueType(f)
}

func printValueType(v any) {
	fmt.Printf("%T\n", v)
	fmt.Printf("%x\n", v)
}
