package main

import "fmt"

func main() {
	s := []int{1, 2, 2, 3}
	removeFirst(s)
	fmt.Println(s)
	// [1 2 3]

	s = []int{1, 2, 2, 3, 2}
	fmt.Println(remove(s, 2))
	s = []int{1, 2, 2, 3}
	fmt.Println(remove(s, 3))
	s = []int{1, 2, 2, 3}
	fmt.Println(remove(s, 1))
	s = []int{1, 2, 2, 3}
	fmt.Println(remove(s, 4))
}

func removeFirst(s []int) {
	s = s[1:]
	fmt.Println(s)
	// [2 3]
	// 此处的局部变量 s 的 len 改变了
}

func remove(s []int, pivot int) []int {
	for index := 0; index < len(s); index++ {
		if s[index] == pivot {
			if len(s)-1 > index {
				s = append(s[:index], s[index+1:]...)
			} else {
				s = s[:index]
			}
			index--
		}
	}
	return s
}
