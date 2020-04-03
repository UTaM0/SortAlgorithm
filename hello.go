package main

import (
	"fmt"
	"sort"
	)

// Sort Array
func main() {
	var N int
	fmt.Scan(&N)
	Hi := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&Hi[i])
	}
	sort.Ints(Hi)
	fmt.Println(Hi)
}
