package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().UTC().Format("2006-01-02T15:04:05Z"))
}
