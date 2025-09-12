package main

import (
	"fmt"
)

// func count(a int) {

// 	if a > 3 {
// 		return
// 	}
// 	fmt.Println(a)
// 	count(a + 1)
// }
// func main() {
// 	count(0)
// }

// func print(i, n int) {
// 	if i > n {
// 		return
// 	}
// 	fmt.Println("Aryan")
// 	print(i+1, n)

// }

// func main() {
// 	print(1, 5)
// }

func oneton(i, n int) {

	if i > n {
		return
	}
	fmt.Println(n)
	oneton(i, n-1)

}

func main() {
	oneton(1, 5)
}
