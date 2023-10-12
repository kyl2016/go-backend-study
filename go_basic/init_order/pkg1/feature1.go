package pkg1

import "fmt"

const Name = "Bob"

func getAge() int {
	fmt.Println("pkg1")
	return 18
}

var age = getAge()

func init() {
	fmt.Println(Name, age)
}
