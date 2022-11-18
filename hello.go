package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world.")
	foo()
	fmt.Println("Back in hello world.")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println("Exiting.")
}

func foo() {
	fmt.Println("Here's another line to run, in foo.")
}

// Control flow:
// (1) sequential
// (2) loop, iterative
// (3) conditional
