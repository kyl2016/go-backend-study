package main

import (
	"fmt"
	"sync"
)

func main() {
	attended := map[string]bool{
		"Ann": true,
		"Joe": true,
	}

	if attended["Ann"] {
		fmt.Println("contains Ann")
	}

	if !attended["Bob"] {
		fmt.Println("not contains Bob")
	}

	person := "Ann"

	if attended[person] {
		fmt.Println(person, "was at the meeting")
	}

	mapping := sync.Map{}
	mapping2 := &sync.Map{}
	mapping.Store(1, mapping2)
	arr := []int{1, 2}
	mapping2.Store(2, &arr)

	v, _ := mapping.Load(1)
	_mapping := v.(*sync.Map)
	v2, _ := _mapping.Load(2)
	fmt.Println(v2)
	ints := v2.(*[]int)
	*ints = append(*ints, 3)

	v, _ = mapping.Load(1)
	_mapping = v.(*sync.Map)
	v2, _ = _mapping.Load(2)
	fmt.Println(v2)

}
