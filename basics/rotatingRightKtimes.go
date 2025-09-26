package main

import "fmt"

func beversingarray(a []int) []int {
	start := 0
	end := len(a) - 1

	for start < end {
		a[start], a[end] = a[end], a[start]
		start++
		end--
	}
	return a

}

func rotatingright(array []int, k int) []int {
	array = beversingarray(array)
	temparr := array[:k]

	temparr = beversingarray(temparr)
	array = array[k:]
	array = beversingarray(array)
	temparr = append(temparr, array...)
	return temparr
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	k := 2

	fmt.Print(rotatingright(arr, k))

}
