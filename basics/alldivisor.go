package main

import (
	"fmt"
	"math"
	"slices"
)

// func alldivisiable(a int) []int {

// 	var f []int

// 	for i := 1; i <= a; i++ {
// 		if a%i == 0 {
// 			f = append(f, i)
// 		}

// 	}
// 	return f

// }

func alldivisiable(a int) []int {

	squareroot := math.Sqrt(float64(a))
	var remainder []int
	var quotient []int
	for i := 1; i <= int(squareroot); i++ {
		if a%i == 0 {
			remainder = append(remainder, i)
			f := a / i
			quotient = append(quotient, f)

		}

	}

	remainder = append(remainder, quotient...)
	slices.Sort(remainder)

	return remainder

}

func main() {
	fmt.Println(alldivisiable(36))
}
