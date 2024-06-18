package main

import (
	"fmt"
	"sync"
)

func PrintSyncMap(cm *sync.Map) {
	cm.Range(func(key, value interface{}) bool {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
		return true
	})
}

func taskN7b() {
	var m sync.Map

	m.Store("foo", "bar")
	m.Store("bar", "baz")
	m.Store("baz", "foo")

	value, ok := m.Load("foo")
	if ok {
		fmt.Println("Value:", value)
	}

	PrintSyncMap(&m)

	m.Delete("foo")

	PrintSyncMap(&m)
}
