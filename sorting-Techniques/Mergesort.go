package main

import "fmt"

func merge(a, b []int) []int {
	var i, j int
	result := []int{}

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}

	for i < len(a) {
		result = append(result, a[i])
		i++
	}
	for j < len(b) {
		result = append(result, b[j])
		j++

	}

	return result
}

func spliting(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	mid := len(a) / 2

	left := spliting(a[:mid])
	right := spliting(a[mid:])

	return merge(left, right)
}

func main() {

	arr := []int{5, 4, 3, 2, 1, 0}

	fmt.Print(spliting(arr))

}

// func main() {

// 	arr1 := []int{1, 3, 6, 10, 15}
// 	arr2 := []int{5, 25, 64}

// 	result := []int{}
// 	var i, j int

// 	for i = range len(arr1) {
// 		if arr1[i] < arr2[j] {
// 			result = append(result, arr1[i])
// 			i++
// 		} else {
// 			result = append(result, arr2[j])
// 			j++
// 		}

// 	}

// 	for i < len(arr1) {
// 		result = append(result, arr1[i])
// 		i++
// 	}
// 	for j < len(arr2) {
// 		result = append(result, arr2[j])
// 		j++
// 	}

// 	fmt.Println(result)
// }

// func main() {
// 	arr1 := []int{10, 15, 3, 6}
// 	arr2 := []int{5, 25, 64, 1}

// 	arrtemp := []int{}

// 	arrtemp = append(arrtemp, arr1...)
// 	arrtemp = append(arrtemp, arr2...)

// 	for x := range len(arrtemp) {
// 		for j := x; j < len(arrtemp); j++ {
// 			if arrtemp[x] > arrtemp[j] {
// 				temp := arrtemp[x]
// 				arrtemp[x] = arrtemp[j]
// 				arrtemp[j] = temp

// 			}

// 		}
// 	}

// 	fmt.Print(arrtemp)

// }
