package main

import "fmt"

func main() {
	var m map[string]int

	if m == nil {
		fmt.Println("is nil")
	}
	m = map[string]int{}
	if m == nil {
		fmt.Println("is nil")
	}
	m["sdf"] = 1
	fmt.Println(len(m))
}
