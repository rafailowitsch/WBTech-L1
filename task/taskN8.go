package task

import "fmt"

// функция SetBit устанавливает i-й бит числа n в 1
func SetBit(n int, i uint) int {
	return n | (1 << i)
}

// функция ClearBit сбрасывает i-й бит числа n в 0
func ClearBit(n int, i uint) int {
	return n &^ (1 << i)
}

// функция taskN8 демонстрирует использование SetBit и ClearBit для установки и сброса битов
func TaskN8() {
	fmt.Print("Initial number: ")
	n := 5 // 101 в двоичном виде
	fmt.Printf("%d in binary is %b\n", n, n)

	fmt.Print("After set first bit: ")
	n = SetBit(n, 1)
	fmt.Printf("%d in binary is %b\n", n, n)

	fmt.Print("After clear zero bit: ")
	n = ClearBit(n, 0)
	fmt.Printf("%d in binary is %b\n", n, n)
}
