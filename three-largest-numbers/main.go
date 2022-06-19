package main

import (
	"fmt"
	"math"
)

func Solution(array []int) [3]int {
  currentLargest := [3]int{math.MinInt32, math.MinInt32, math.MinInt32}
  for _, val := range array {
    if val > currentLargest[2] {
      appendToArray(&currentLargest, 2, val)
    } else if val > currentLargest[1] {
      appendToArray(&currentLargest, 1, val)
    } else {
      appendToArray(&currentLargest, 0, val)
    }
  }
  return currentLargest
}

func appendToArray(array *[3]int, index, value int) {
  for i := 0; i <= index; i++ {
    if i == index {
      array[i] = value
    } else {
      array[i] = array[i+1]
    }
  }
}

func main() {
  array := []int{141,1,17,-7,-17,-27,18,541,8,7,7}
  fmt.Printf("Result => %d\n", Solution(array))
}
