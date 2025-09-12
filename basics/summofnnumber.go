package main

import "fmt"

var b int

func sumofn(i int) int {

	if i == 0 {
		return 0
	}
	return i + sumofn(i-1)

}

func main() {
	fmt.Println(sumofn(5))
}
