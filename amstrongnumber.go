package main

import (
	"fmt"
	"math"
)

func amstrong(num int) bool {

	var temp int
	var k int
	temp = num

	for temp > 0 {
		temp = temp / 10
		k++

	}
	var input int
	input = num
	var final int

	for input > 0 {
		a := input % 10
		input = input / 10
		final = final + int(math.Pow(float64(a), float64(k)))

	}

	return num == final
}

func main() {

	fmt.Println(amstrong(153))

}
