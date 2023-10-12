package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines, err := ReadLines("./go.mod")
	fmt.Println(lines, err)
	return
}

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var r []string
	for scanner.Scan() {
		r = append(r, scanner.Text())
	}
	return r, nil
}
