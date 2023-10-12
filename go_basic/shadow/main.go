package main

import "fmt"

func main() {
	v := "outer"
	fmt.Println(v)
	{
		v := "inner"
		fmt.Println(v)
		{
			fmt.Println(v)
		}
	}
	{
		fmt.Println(v)
	}
	if true {
		v := 1
		fmt.Println(v)
	}
	fmt.Println(v)
}
