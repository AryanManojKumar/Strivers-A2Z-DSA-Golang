package main

import "fmt"

func fact(a int) int {
	if a == 1 {
		return 1
	}
	return a * fact(a-1)

}

func main() {
	fmt.Println(fact(5))
}
