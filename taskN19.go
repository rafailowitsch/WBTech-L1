package main

import "fmt"

func reverseSrting(str string) string {
	runes := []rune(str)

	l := len(runes)
	mid := l / 2

	for i := 0; i < mid; i++ {
		runes[i], runes[l-1-i] = runes[l-1-i], runes[i]
	}

	str = string(runes)
	return str
}

func taskN18() {
	str := "главрыба"
	fmt.Println("String before reverseStr: ", str)

	str = reverseSrting(str)
	fmt.Println("String afer reverseStr: ", str)
}
