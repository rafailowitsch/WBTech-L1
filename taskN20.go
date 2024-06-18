package main

import (
	"fmt"
	"strings"
)

func reverseWords(str string) string {
	words := strings.Fields(str)

	l := len(words)
	mid := l / 2
	println(mid)

	for i := 0; i < mid; i++ {
		words[i], words[l-1-i] = words[l-1-i], words[i]
		fmt.Println(i, words[i], words[l-1-i])
	}

	str = strings.Join(words, " ")

	return str
}

func taskN19() {
	str := "snow dog sun moon"
	fmt.Println("String before reverse words: ", str)

	str = reverseWords(str)
	fmt.Println("String after reverse words:", str)
}
