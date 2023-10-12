package main

import "fmt"

var f func() {
	fmt.Println("f")
}

func init() {
	fmt.Println("init", f)
}

func main() {
	f()
}
