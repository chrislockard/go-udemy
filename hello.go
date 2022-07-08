package main

import "fmt"

func main() { //entry point
	fmt.Println("Hello, world!")
	foo()

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// Statements are made up of expressions
	x := 42
	// Statements are made up of expressions
	fmt.Println(x)
	// Let's print the types of values
	printReturnTypes(x)
	greet := "Hello!"
	printReturnTypes(greet)
	f := 3.14159
	printReturnTypes(f)
}

func foo() {
	fmt.Println("I'm in foo")
}

func printReturnTypes(v any) {
	fmt.Printf("%T\n", v)
	fmt.Printf("%x\n", v)
}

// control flow:
// 1 sequence
// 2 loop; iterative
// 3 conditional
