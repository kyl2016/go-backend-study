package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	handler()
}

func handler() {
	log()
}

func log() {
	fmt.Println(string(debug.Stack()))
}
