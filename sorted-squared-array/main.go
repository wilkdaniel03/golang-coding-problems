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

func main() {
	input := []int{1, 2, 3, 5, 6, 8, 9}
	fmt.Printf("First solution's result => %d\n", sortAlgoSolution(input))
}
