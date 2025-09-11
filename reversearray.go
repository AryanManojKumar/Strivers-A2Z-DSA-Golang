package main

import "fmt"

func reversingarray(a []int, n int) []int {
	var b []int
	fmt.Println(n)
	for i := a[n-1]; i >= 0; i-- {
		b = append(b, a[i])
	}
	return b

}

func main() {
	fmt.Println(reversingarray([]int{1, 2, 3, 4}, 4))
}
