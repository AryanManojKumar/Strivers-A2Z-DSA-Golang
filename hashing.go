package main

import (
	"fmt"
)

func main() {
	var arr = []int{1, 3, 2, 1, 3}
	quar := []int{1, 4, 2, 3, 12}

	storingmap := make(map[int]int)

	for _, key := range arr {

		storingmap[key] = storingmap[key] + 1

	}

	for _, quer := range quar {
		fmt.Println(storingmap[quer])
	}

}
