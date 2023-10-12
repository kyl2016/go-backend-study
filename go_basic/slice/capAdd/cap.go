package main

import "fmt"

func main() {
	s := make([]int, 0, 3)
	fmt.Println(len(s), cap(s))
	s = append(s, 1)
	fmt.Println(s)

	s2 := make([]int, 1, 3)
	fmt.Println(len(s2), cap(s2))
	s2 = append(s2, 1)
	fmt.Println(s2)
}
