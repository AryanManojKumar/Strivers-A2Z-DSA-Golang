package main

import "fmt"

func shorting(a []int) []int {
	var min int

	for i := range a {
		min = i
		for f := i + 1; f < len(a); f++ {
			if a[i] > a[f] {
				min = f
			}
			b := a[i]
			a[i] = a[min]
			a[min] = b
		}

	}
	return a

}

func main() {
	fmt.Println(shorting([]int{13, 46, 24, 52, 20, 9}))
}
