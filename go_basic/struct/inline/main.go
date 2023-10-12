package main

import (
	"context"
	"fmt"
)

func main() {
	// have to initialize the embeded struct as a whole
	// s := Student{
	// 	Name:  "Bob",
	// 	Class: "一班",
	// }
	s := Student{
		People: People{Name: "Hello"},
		Class:  "二班",
	}
	fmt.Println(s)

	var s2 Student
	s2.Name = "World"
	s2.Class = "三班"
	fmt.Println(s2)

	// Embedding structs also works with methods
	s2.Print()

	// 参考：context.timerCtx
}

type Student struct {
	People
	Class string
	xxx
}

type People struct {
	Name string
}

func (p People) Print() {
	fmt.Println(p.Name)
}

type 