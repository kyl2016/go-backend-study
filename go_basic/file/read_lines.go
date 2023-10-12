package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rw, err := os.Open("./basic.go")
	if err != nil {
		panic(err)
	}
	defer rw.Close()
	sb := bufio.NewScanner(rw)
	for sb.Scan() {
		//do something
		fmt.Println(sb.Text())
	}
	if err := sb.Err(); err != nil {
		panic(err)
	}
}
