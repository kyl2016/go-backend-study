package main

import "strconv"

func main() {
	a := 1
	b := 2

	r := max(a, b)
	println(`max: ` + strconv.Itoa(r))
}

func max(a int, b int) int {
	if a >= b {
		return a
	}

	return b
}

// go build -gcflags -S ./main.go
