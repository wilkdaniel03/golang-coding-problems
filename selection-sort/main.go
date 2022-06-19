package main

import "fmt"

func Solution(array []int) []int {
  for currentIdx := 0; currentIdx < len(array)-1; currentIdx++ {
    smallestIdx := currentIdx
    for comparedIdx := currentIdx+1; comparedIdx < len(array); comparedIdx++ {
      if array[comparedIdx] < array[smallestIdx] {
        smallestIdx = comparedIdx
      }
    }
    array[currentIdx], array[smallestIdx] = array[smallestIdx], array[currentIdx]
  }
  return array
}

func main() {
  array := []int{8,5,2,9,5,6,3}
  fmt.Printf("Result => %v\n", Solution(array))
}
