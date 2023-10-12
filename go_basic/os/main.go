package main

import (
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("go.mod")
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(data)
}
