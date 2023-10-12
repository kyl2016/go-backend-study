package main

import (
	"fmt"
)

func main() {
	m := map[string]interface{}{
		"age": 1,
	}
	add(m)
	fmt.Println(m)
}

func add(m map[string]interface{}) {
	m2 := m
	m2["name"] = "Bob"
}
