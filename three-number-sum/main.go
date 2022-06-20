package main

import (
	"fmt"
	"sort"
)

func Solution(array []int, target int) [][]int {
  sort.Ints(array)
  triplets := [][]int{}
  for i := 0; i < len(array)-2; i++ {
    start, end := i+1, len(array)-1
    for start < end {
      currentSum := array[i] + array[start] + array[end]
      if currentSum == target {
        triplets = append(triplets, []int{array[i], array[start], array[end]})
        start += 1
        end -= 1
      } else if currentSum < target {
        start += 1
      } else {
        end -= 1
      }
    }
  }
  return triplets
}

func main() {
  array := []int{12,3,1,2,-6,5,-8,6}
  fmt.Printf("Result => %v\n", Solution(array, 0))
}
