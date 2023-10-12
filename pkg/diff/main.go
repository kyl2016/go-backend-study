package main

import (
	"fmt"

	"github.com/r3labs/diff/v2"
)

func main() {
	from := []int{1, 2, 3, 4}
	to := []int{1, 2, 4}

	changelog, _ := diff.Diff(from, to)
	fmt.Printf("%+v\n", changelog)
	// [{Type:delete Path:[2] From:3 To:<nil> parent:<nil>}]

	// 改变元素顺序，不认为是 change
	to = []int{1, 2, 4, 3}
	changelog, _ = diff.Diff(from, to)
	fmt.Printf("%+v\n", changelog)

	a := Order{
		ID:    "1234",
		Items: []int{1, 2, 3, 4},
	}

	b := Order{
		ID:    "1234",
		Items: []int{1, 2, 5},
	}

	changelog, _ = diff.Diff(a, b)
	fmt.Printf("%+v\n", changelog)
	// [{Type:update Path:[items 2] From:3 To:5 parent:<nil>} {Type:delete Path:[items 3] From:4 To:<nil> parent:<nil>}]
}

type Order struct {
	ID    string `diff:"id"`
	Items []int  `diff:"items"`
}
