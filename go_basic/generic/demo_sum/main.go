package main

import (
	"fmt"
)

func main() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	int32s := map[string]int32{
		"first":  35,
		"second": 26,
	}

	fmt.Println(SumIntsOrFloats[string, int64](ints))
	fmt.Println(SumIntsOrFloats(ints))
	fmt.Println(SumIntsOrFloats(int32s))
	fmt.Println(SumIntsOrFloats(floats))
}

func SumIntsOrFloats[K comparable, V int64 | float64 | int32](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
