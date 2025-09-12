package main

import (
	"fmt"
)

func main() {

	n := 5

	for i := 1; i < n; i++ {

		for range n - i {
			fmt.Print(" ")
		}
		for range 2*i - 1 {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := 1; i < n; i++ {

		for f := n; f >= n-i; f-- {
			fmt.Print(" ")
		}

		for j := n; j >= 2*i-1; j-- {
			fmt.Print("*")

		}
		fmt.Println()

	}

}
