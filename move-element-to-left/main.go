package main

import "fmt"

func Solution(array []int, toMove int) []int {
  startIdx, endIdx := 0, len(array)-1
  for startIdx < endIdx {
    for startIdx < endIdx &&
    array[endIdx] == toMove {
      endIdx -= 1
    }
    if array[startIdx] == toMove {
      array[startIdx], array[endIdx] = 
        array[endIdx], array[startIdx]
    }
    startIdx += 1
  }
  return array
}

func main() {
  array := []int{2,1,2,2,2,3,4,2}
  fmt.Printf("Result => %v\n", Solution(array, 2))
}
