package main

import "fmt"

func SetBit(n int, i uint) int {
	return n | (1 << i)
}

func ClearBit(n int, i uint) int {
	return n &^ (1 << i)
}

func taskN8() {

	fmt.Print("Initial number: ")
	n := 5 // 101
	fmt.Printf("%d in binary is %b\n", n, n)

	fmt.Print("After set first bit: ")
	n = SetBit(n, 1)
	fmt.Printf("%d in binary is %b\n", n, n)

	fmt.Print("After clear zero bit: ")
	n = ClearBit(n, 0)
	fmt.Printf("%d in binary is %b\n", n, n)
}
