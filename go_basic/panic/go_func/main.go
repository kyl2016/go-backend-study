package main

// 主线程会 panic
func main() {
	go func() {
		panic("test")
	}()

	select {}
}
