package main

import (
	"fmt"
	"sort"
)

func Solution(
  redShirtSpeeds []int,
  blueShirtSpeeds []int,
  fastest bool,
) int {
  sort.Sort(sort.Reverse(sort.IntSlice(redShirtSpeeds)))
  if fastest {
    sort.Ints(blueShirtSpeeds)
  } else {
    sort.Sort(sort.Reverse(sort.IntSlice(blueShirtSpeeds)))
  }
  currentSpeed := 0
  for i := 0; i < len(redShirtSpeeds); i++ {
    redCyclist, blueCyclist := redShirtSpeeds[i], blueShirtSpeeds[i]
    currentSpeed += max(redCyclist, blueCyclist)
  }
  return currentSpeed
}

func max(a int, b int) int {
  if a > b {
    return a
  }
  return b
}

func main() {
  redShirtSpeeds := []int{5,5,3,9,2}
  blueShirtSpeeds := []int{3,6,7,2,1}
  fmt.Printf("Result => %v\n", Solution(redShirtSpeeds, blueShirtSpeeds, true))
}
