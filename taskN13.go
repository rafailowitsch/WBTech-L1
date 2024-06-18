package main

import "fmt"

func taskN13() {
	// первый способ
	a := 5
	b := 2

	fmt.Printf("a = %d, b = %d\n", a, b)

	a = a + b // = 5 + 2 = 7
	b = a - b // = 7 - 2 = 5

	fmt.Printf("a = %d, b = %d\n", a, b)

	// второй способ
	fmt.Printf("a = %d, b = %d\n", a, b)

	a, b = b, a

	fmt.Printf("a = %d, b = %d\n", a, b)
}
