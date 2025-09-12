package main

import (
	"fmt"
)

func reversing(x int) int {
	negLimit := -1 << 31
	posLimit := (1 << 31) - 1
	var b int

	for x != 0 {
		var a int

		a = x % 10
		x = x / 10

		if b > posLimit/10 || (b == posLimit/10 && a > 7) {
			return 0
		}
		if b < negLimit/10 || (b == negLimit/10 && a < -8) {
			return 0
		}

		b = b*10 + a

	}

	return b
}

func main() {
	fmt.Println(reversing(-1234))
}
