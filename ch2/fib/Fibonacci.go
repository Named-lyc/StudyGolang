package main

import "fmt"

func main() {
	var x = 0
	x = fib(10)
	fmt.Println(x)
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
