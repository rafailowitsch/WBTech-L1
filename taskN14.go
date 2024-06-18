package main

import "fmt"

func taskN14() {
	sl := []interface{}{5, "string", true, make(chan interface{}), struct{}{}}

	for i, elm := range sl {
		eType := getType(elm)
		fmt.Printf("Elem [%d] has type %s\n", i, eType)
	}
}

func getType(iface interface{}) string {
	switch iface.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan interface{}:
		return "chan of interface{}"
	default:
		return "unknown type"
	}
}
