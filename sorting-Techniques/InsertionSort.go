package main

import (
	"fmt"
)

func inserstionsort(a []int) []int {

	for i := 1; i < len(a); i++ {
		temp := a[i]

		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				a[j] = a[j-1]
				a[j-1] = temp
			}
		}
	}
	return a

}

func main() {
	fmt.Println(inserstionsort([]int{13, 46, 24, 52, 20, 9}))
}

// i like to first do it in main to dry run in my mind

// func main() {
// 	arr := []int{10, 3, 8, 1, 5}

// 	for i := 1; i < len(arr); i++ {
// 		key := arr[i]

// 		for j := i; j > 0; j-- {

// 			if arr[j] < arr[j-1] {
// 				arr[j] = arr[j-1]
// 				arr[j-1] = key
// 			}

// 		}

// 	}
// 	fmt.Println(arr)
// }
