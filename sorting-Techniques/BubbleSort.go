package main

import "fmt"

func swap(a int, b int) (int, int) {

	temp := a
	a = b
	b = temp
	return a, b
}

func bubblesort(arr []int) []int {

	for i := len(arr) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = swap(arr[j], arr[j+1])
			}
		}

	}

	return arr
}

func main() {
	fmt.Println(bubblesort([]int{7, 5, 9, 2, 8}))

}

// i like to first do it in main to dry run in my mind
// func main() {

// 	arr := []int{5, 4, 3, 2, 1}

// 	for pass := 0; pass < len(arr); pass++ {

// 		for i := 0; i < len(arr)-1-pass; i++ {

// 			if arr[i] > arr[i+1] {
// 				fmt.Println(arr[i], arr[i+1])

// 				temp := arr[i]
// 				arr[i] = arr[i+1]
// 				arr[i+1] = temp

// 			}
// 		}
// 	}

// 	fmt.Println(arr)
// }
