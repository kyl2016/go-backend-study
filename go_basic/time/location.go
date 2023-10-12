package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now.UTC().Format(time.RFC3339))

	location, err := time.LoadLocation("GMT0")
	if err != nil {
		panic(err)
	}
	fmt.Println(now.In(location).Format(time.RFC3339))

	location, err = time.LoadLocation("Local")
	if err != nil {
		panic(err)
	}
	fmt.Println(now.In(location).Format(time.RFC3339))

	location, err = time.LoadLocation("Europe/Rome")
	if err != nil {
		panic(err)
	}
	fmt.Println(now.In(location).Format(time.RFC3339))
}
