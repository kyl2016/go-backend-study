package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	fmt.Println(now.Format("2006-01-02 15:04:05.999"))
	fmt.Println(now.Format("2006-01-02 15:04:5.999"))

	// 时区为 0
	t, err := time.Parse(time.RFC3339, "2020-12-02T03:30:30+00:00")
	fmt.Println(t, err)
	fmt.Println(t.Format(time.RFC3339))

	// 时区为东八区
	t, err = time.Parse(time.RFC3339, "2020-12-02T03:30:30+08:00")
	fmt.Println(t, err)
	fmt.Println(t.Format(time.RFC3339))

	fmt.Println(now.GoString())
	fmt.Println(time.UnixMilli(now.UnixMilli()).UnixMilli())
	fmt.Println(now.UnixMilli())
}
