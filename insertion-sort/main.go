package main

import "fmt"

func Solution(array []int) []int {
  for i := range array {
    for j := i; j > 0 && array[j] < array[j-1]; j-- {
      array[j], array[j-1] = array[j-1], array[j]
    }
  }
  return array
}

func main() {
  array := []int{8,5,2,9,5,6,3}
  fmt.Printf("Result => %v\n", Solution(array))
}
