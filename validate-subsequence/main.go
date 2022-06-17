package main

import "fmt"

func firstSolution(array []int, sequence []int) bool {
	arrayIdx, sequenceIdx := 0, 0
	for arrayIdx < len(array) && sequenceIdx < len(sequence) {
		if array[arrayIdx] == sequence[sequenceIdx] {
			sequenceIdx += 1
		}
		arrayIdx += 1
	}
	return sequenceIdx == len(sequence)
}

func main() {
	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{1, 6, -1, 10}
	fmt.Printf("First solution's result => %v\n", firstSolution(array, sequence))
}
