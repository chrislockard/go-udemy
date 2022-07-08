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
}

func foo() {
	fmt.Println("I'm in foo")
}

// control flow:
// 1 sequence
// 2 loop; iterative
// 3 conditional
