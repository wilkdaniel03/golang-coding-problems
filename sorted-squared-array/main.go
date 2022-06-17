package main

import (
	"fmt"
	"sort"
)

func sortAlgoSolution(array []int) []int {
	squaredArray := make([]int, len(array))
	for idx, num := range array {
		squaredArray[idx] = num * num
	}
	sort.Ints(squaredArray)
	return squaredArray
}

func putInPlaceSolution(array []int) []int {
	squaredArray := make([]int, len(array))
	start, end := 0, len(array)-1
	for i := len(array) - 1; i >= 0; i-- {
		if abs(array[start]) > abs(array[end]) {
			squaredArray[i] = array[start] * array[start]
			start += 1
		} else {
			squaredArray[i] = array[end] * array[end]
			end -= 1
		}
	}
	return squaredArray
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	input := []int{1, 2, 3, 5, 6, 8, 9}
	fmt.Printf("First solution's result => %d\n", sortAlgoSolution(input))
	fmt.Printf("First solution's result => %d\n", putInPlaceSolution(input))
}
