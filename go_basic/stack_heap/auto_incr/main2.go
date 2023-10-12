package main

func main() {
	var x [10]int
	a(x)
}

//go:noinline
func a(x [10]int) {
	println(`func a`)
	var y [100]int
	b(y)
}

//go:noinline
func b(x [100]int) {
	println(`func b`)
	var y [1000]int
	c(y)
}

//go:noinline
func c(x [1000]int) {
	println(`func c`)
}
