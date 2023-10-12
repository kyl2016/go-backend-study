package main

import "fmt"

func main() {
	sample1()
	sample2()
}

func sample2() {
	arr := []map[string]int{}
	arr = append(arr, map[string]int{"a:": 1, "b": 2, "c": 3})

	var arrI interface{}
	arrI = arr

	var i interface{}
	i = map[string]interface{}{
		"key": arrI,
	}

	r, ok := i.(map[string]interface{})
	fmt.Println(r, ok)

	query(r)
}

func query(m map[string]interface{}) {
	v := m["key"]
	r, ok := v.([]map[string]int)
	fmt.Println(r, ok)
}

func sample1() {
	var i interface{}
	arr := []map[string]int{}
	arr = append(arr, map[string]int{"a:": 1, "b": 2, "c": 3})
	i = arr

	s, ok := i.([]map[string]int)
	if ok {
		for _, it := range s {
			for k, v := range it {
				fmt.Println(k, v)
			}
		}
	}
	fmt.Println(s, ok)
}
