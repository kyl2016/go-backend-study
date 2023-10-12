package main

import "fmt"

func main() {
	fmt.Println(isValidCUID("000-1"))
	fmt.Println(isValidCUID("000-00-00"))
	fmt.Println(isValidCUID(""))
}

func isValidCUID(cuid string) bool {
	for _, ch := range cuid {
		if ch != '0' && ch != '-' {
			return true
		}
	}
	return false
}
