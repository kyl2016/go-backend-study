package main

import "fmt"

type Person struct {
	Name string
}

func main() {
	p := Person{Name: "123"}
	change(&p)
	fmt.Println(p)
	// {123}
	change2(&p)
	fmt.Println(p)
	// {456}
}

func change(p *Person) {
	p = &Person{Name: "456"}
}

func change2(p *Person) {
	*p = Person{Name: "456"}
}
