package main

import (
	"fmt"
	"strings"
)

var justString string

func taskN15() {
	someFunc()
	fmt.Println(justString)
}

func createHugeString(size int) string {
	repeatedString := strings.Repeat("a", size)
	return repeatedString
}

func someFunc() {
	v := createHugeString(1 << 4)

	justString = string([]rune(v)[:5])
}
