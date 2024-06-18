package main

import "fmt"

func taskN13() {
	a := 5
	b := 2

	fmt.Printf("a = %d, b = %d\n", a, b)

	a = a + b // = 5 + 2 = 7
	b = a - b // = 7 - 2 = 5

	fmt.Printf("a = %d, b = %d\n", a, b)
}
