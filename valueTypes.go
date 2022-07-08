package main

import "fmt"

func main() { //entry point
	x := 42

	// Let's print the types of values
	printReturnTypes(x)
	greet := "Hello!"
	printReturnTypes(greet)
	f := 3.14159
	printReturnTypes(f)
}

func printReturnTypes(v any) {
	fmt.Printf("%T\n", v)
	fmt.Printf("%x\n", v)
}
