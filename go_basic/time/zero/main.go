package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Time{}
	fmt.Println("t.IsZero:", t.IsZero())

	var t2 time.Time
	fmt.Println("t2.IsZero:", t2.IsZero())

	if !t.IsZero() && t.After(time.Now()) {
		fmt.Println("ok")
	}
}
