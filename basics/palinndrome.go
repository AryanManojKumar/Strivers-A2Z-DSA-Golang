package main

import "fmt"

func palindrome(n int) bool {
	go func() bool {
		if n < 0 {
			return false
		}
	}()

	input := n
	var reverse int

	for n > 0 {
		k := n % 10
		n = n / 10
		reverse = reverse*10 + k

	}
	return input == reverse
}

func main() {
	fmt.Println(palindrome(141))
}
