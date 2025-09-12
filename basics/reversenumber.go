package main

import "fmt"

func reverse(n int) int {
	var b int

	for n > 0 {
		var a int

		a = n % 10
		n = n / 10

		b = b*10 + a

	}
	return b

}

func main() {
	fmt.Println(reverse(654321))
}
