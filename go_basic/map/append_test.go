package main

import (
	"fmt"
	"testing"
)

func TestMap_Append(t *testing.T) {
	var m2 map[int]int
	fmt.Println(m2[1])

	m := map[int]int{}
	add22(m, 1, 1)
	fmt.Println(m)
	fmt.Println("")

	add22(m, 2, 2)
	fmt.Println(m)

	remove(m, 2)
	fmt.Println(m)
	fmt.Println("ok")
}

func add22(m map[int]int, k, v int) {
	m[k] = v
}

func remove(m map[int]int, k int) {
	delete(m, k)
}
