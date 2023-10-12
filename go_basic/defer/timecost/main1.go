package main

import (
	"fmt"
	"time"
)

func timeCost() func() {
	now := time.Now()
	return func() {
		fmt.Println("timeCost")
		fmt.Println("耗时", time.Since(now))
	}
}

func do() {
	defer timeCost()()
	time.Sleep(time.Second)

	return
}

func main() {
	do()
}
