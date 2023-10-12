package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("error info")
	fmt.Printf("%s\n", err)
	fmt.Printf("%v\n", err)
	fmt.Printf("%+v\n", err)
}
