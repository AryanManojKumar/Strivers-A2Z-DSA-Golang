package main

import (
	"fmt"
)

var memory = map[int]int{}

func fib(n int) int {

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if val, inside := memory[n]; inside {
		return val
	}
	memory[n] = fib(n-1) + fib(n-2)
	return memory[n]

}

func main() {
	fmt.Println(fib(50))

}
