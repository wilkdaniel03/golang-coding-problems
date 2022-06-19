package main

import "fmt"

type SpecialArray []interface{}

func Solution(array SpecialArray) int {
  return helper(array, 1)
}

func helper(array SpecialArray, depth int) int {
  sum := 0
  for _, val := range array {
    if v, ok := val.(SpecialArray); ok {
      sum += helper(v, depth+1)
    } else if v, ok := val.(int); ok {
      sum += v
    }
  }
  return sum * depth
}

func main() {
  array := SpecialArray{5,2,SpecialArray{7,-1},3,SpecialArray{6,SpecialArray{-13,8},4}}
  fmt.Printf("Result => %d\n", Solution(array))
}
