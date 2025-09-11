package main

import "fmt"

func calculatingfreq(arr []int) {

	f := make(map[int]int)

	for _, key := range arr {
		f[key] = f[key] + 1

	}

	var r int
	var hk int

	var lowest int
	var lk int

	fmt.Println(f)

	for keyys, frequese := range f {

		if r < frequese {
			r = frequese
			hk = keyys
			lowest = frequese
		}

		if lowest > frequese {
			fmt.Println(lowest)

			lk = keyys
			lowest = frequese

		}

	}

	fmt.Println("the highest one is ", hk)
	fmt.Println("lowest one is", lk)

}

func main() {

	calculatingfreq([]int{2, 2, 3, 4, 4, 2})

}
