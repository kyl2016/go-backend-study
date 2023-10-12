package main

import (
	"fmt"
	"sync"
)

// append 并发不安全，应该用 channel
func main2() {
	for j := 0; j < 1000; j++ {
		var arr []int
		wg := sync.WaitGroup{}
		count := 0
		for i := 0; i < 3; i++ {
			wg.Add(1)
			go func(i int) {
				arr = append(arr, i)
				count++
				wg.Done()
			}(i)
		}
		wg.Wait()

		if count != 3 {
			panic(count)
		}

		if len(arr) != 3 {
			fmt.Println(arr)
			panic(arr)
		}
	}

}
