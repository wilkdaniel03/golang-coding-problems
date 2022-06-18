package main

import (
	"fmt"
	"sort"
)

func Solution(
  redShirtStudents []int,
  blueShirtStudents []int,
) bool {
  sort.Sort(sort.Reverse(sort.IntSlice(redShirtStudents)))
  sort.Sort(sort.Reverse(sort.IntSlice(blueShirtStudents)))
  firstRow := "RED"
  if redShirtStudents[0] > blueShirtStudents[0] {
    firstRow = "BLUE"
  }
  for i := 0; i < len(redShirtStudents); i++ {
    redStudent, blueStudent := redShirtStudents[i], blueShirtStudents[i]
    if firstRow == "RED" {
      if redStudent >= blueStudent {
        return false
      }
    } else {
      if blueStudent >= redStudent {
        return false
      }
    }
  }
  return true
}

func main() {
  redShirtHeights := []int{5,8,1,3,4}
  blueShirtHeights := []int{6,9,2,4,5}
  fmt.Printf("Result => %v\n", Solution(redShirtHeights, blueShirtHeights))
}
