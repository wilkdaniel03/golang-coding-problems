package main

import "fmt"

func nestedLoopSolution(array []int, targetSum int) []int {
  for i := 0; i < len(array); i++ {
    firstNum := array[i]
    for j := i+1; j < len(array); j++ {
      secondNum := array[j]
      if firstNum + secondNum == targetSum {
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

func main() {
  input := []int{3,5,-4,8,11,1,-1,6}
  fmt.Printf("First solution's result => %d\n", nestedLoopSolution(input, 10))
  fmt.Printf("Second solution's result => %d\n", hashMapSolution(input, 10))
}
