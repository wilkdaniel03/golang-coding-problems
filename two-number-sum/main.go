package main

import "fmt"

func nestedLoopSolution(array []int, targetSum int) []int {
	for i := 0; i < len(array); i++ {
		firstNum := array[i]
		for j := i + 1; j < len(array); j++ {
			secondNum := array[j]
			if firstNum+secondNum == targetSum {
				return []int{firstNum, secondNum}
			}
		}
	}
	return []int{}
}

func hashMapSolution(array []int, targetSum int) []int {
	nums := map[int]bool{}
	for _, num := range array {
		possibleMatch := targetSum - num
		if _, found := nums[possibleMatch]; found {
			return []int{num, possibleMatch}
		}
		nums[num] = true
	}
	return []int{}
}

func pointersSolution(array []int, targetSum int) []int {
	smallestIdx, largestIdx := 0, len(array)-1
	for smallestIdx < largestIdx && largestIdx < len(array) {
		currentResult := array[smallestIdx] + array[largestIdx]
		if currentResult < targetSum {
			smallestIdx += 1
		} else if currentResult > targetSum {
			largestIdx -= 1
		} else {
			return []int{array[smallestIdx], array[largestIdx]}
		}
	}
	return []int{}
}

func main() {
	input := []int{3, 5, -4, 8, 11, 1, -1, 6}
	fmt.Printf("First solution's result => %d\n", nestedLoopSolution(input, 10))
	fmt.Printf("Second solution's result => %d\n", hashMapSolution(input, 10))
	fmt.Printf("Third solution's result => %d\n", pointersSolution(input, 10))
}
