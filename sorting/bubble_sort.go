package main

import "fmt"

func bubbleSort(numbers []int) {
	for i := 0; i < len(numbers); i++ {
		if !sweep(numbers, i) {
			return
		}
	}
}

func sweep(numbers []int, currentPass int) bool {
	var idx int
	idxNext := idx + 1
	n := len(numbers)
	var swap bool

	for idxNext < (n - currentPass) {
		a := numbers[idx]
		b := numbers[idxNext]

		if a > b {
			numbers[idx] = b
			numbers[idxNext] = a
			swap = true
		}
		idx++
		idxNext = idx + 1
	}

	return swap
}

func main() {
	org := []int{1, 3, 2, 4, 8, 6, 7, 2, 3, 0}
	fmt.Println(org)

	bubbleSort(org)
	fmt.Println(org)
}
