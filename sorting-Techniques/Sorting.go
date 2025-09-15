package main

import (
	"fmt"
)

func main() {

	arr := []int{13, 46, 24, 52, 20, 9}

	for i := range len(arr) {

		for j := i; j < len(arr); j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp

			}

		}

	}

	fmt.Println(arr)

}
