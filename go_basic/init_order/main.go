package main

import (
	"fmt"

	_ "github.com/kyl2016/Play-With-Golang/basic/init_order/pkg1"
	_ "github.com/kyl2016/Play-With-Golang/basic/init_order/pkg2"
)

const Name = "Kitty"

var age = 20

func init() {
	fmt.Println(Name, age)
}

func main() {

	fmt.Println("main")
}
