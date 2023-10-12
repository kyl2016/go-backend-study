package main

import "fmt"

func main() {
	arr := make([]int, 1, 10)
	fmt.Println(cap(arr))
	fmt.Printf("%p\n", &arr)

	addMore(arr)
	fmt.Println(len(arr))
	fmt.Println(arr)
	fmt.Printf("%p\n", &arr)

	fmt.Println(arr[0:2])

	modify(arr)
	fmt.Printf("%p\n", &arr)

	addMoreUseReference(&arr)
}

func addMore(arr []int) {
	arr = append(arr, 21)
	fmt.Println("addMore ", cap(arr))
	fmt.Printf("addMore %p\n", &arr)
}

func addMoreUseReference(arr *[]int) {
	*arr = append(*arr, 1)
	fmt.Printf("addMore %p\n", arr)
}

func modify(arr []int) {
	arr[0] = 2
	fmt.Printf("modify %p\n", &arr)
}
