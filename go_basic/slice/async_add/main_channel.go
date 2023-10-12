package main

import (
	"fmt"
	"sync"
)

// append 并发不安全，应该用 channel
func main() {
	for j := 0; j < 10000; j++ {
		ch := make(chan int, 3)
		wg := sync.WaitGroup{}
		for i := 0; i < 3; i++ {
			wg.Add(1)
			go func(i int) {
				ch <- i
				wg.Done()
			}(i)
		}
		wg.Wait()
		close(ch)
		if len(ch) != 3 {
			for item := range ch {
				fmt.Println(item)
			}
		}
	}
}
