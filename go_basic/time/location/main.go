package main

import (
	"fmt"
	"time"
)

func main() {
	location, _ := time.LoadLocation("Asia/Shanghai")
	fmt.Println(time.Now().In(location))

	l := time.FixedZone("CST", 3600*8) // 东八区
	fmt.Println(time.Now().In(l))
}
