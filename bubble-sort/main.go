package main

import "fmt"

func Solution(array []int) []int {
  isSorted := false
  for isSorted == false {
    isSorted = true
    for i := 0; i < len(array) - 1; i++ {
      if array[i] > array[i+1] {
        array[i], array[i+1] = array[i+1], array[i]
        isSorted = false
      }
    }
  }
  return array
}

func main() {
  array := []int{8,5,2,9,5,6,3}
  fmt.Printf("Result => %v\n", Solution(array))
}
