package main

import "fmt"

func main() {
	s := []int{1, 1, 2, 3, 4, 1, 2}
	fmt.Println(s)
	fmt.Println(removeDuplicated2(s))
}

func removeDuplicate(s []int) []int {
	fmt.Println(s[1:]) // FIXME why 不崩溃，但 s[1] 会崩溃
	for i := 0; i < len(s); i++ {
		if s[i] == 1 {
			s = append(s[:i], s[i+1:]...)
			i--
		}
	}
	return s
}

func removeDuplicated2(s []int) []int {
	m := map[int]bool{}
	var distinctEvents []int
	for _, item := range s {
		ok := m[item]
		fmt.Println(ok)
		if !ok {
			distinctEvents = append(distinctEvents, item)
			m[item] = true
		}
	}

	return distinctEvents
}
