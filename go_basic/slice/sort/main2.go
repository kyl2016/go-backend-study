package main

import (
	"fmt"
	"sort"
)

func main() {
	eventNames := []string{"b", "c", "a"}
	sort.Strings(eventNames)
	fmt.Println(eventNames)

	sort.SliceStable(eventNames, func(i, j int) bool {
		return eventNames[i] < eventNames[j]
	})
}
