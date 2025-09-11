package main

import (
	"fmt"
	"math"
)

// func findingnemo(a int) int {
// 	var count int

// 	for a > 0 {
// 		a = a / 10
// 		count++

// 	}
// 	return (count)

// }

func logging(n float64) int {
	return int(math.Log10(n) + 1)
}

func main() {
	fmt.Println(logging(1234567))

}
