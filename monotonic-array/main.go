package main

import "fmt"

func FirstSolution(array []int) bool {
  if len(array) <= 2 {
    return true
  }

  direction := array[1] - array[0]
  for i := 2; i < len(array); i++ {
    if direction == 0 {
      direction = array[i] - array[i-1]
      continue
    }
    if checkIfBreaksDirection(direction, array[i], array[i-1]) {
      return false
    }
  }
  return true
}

func checkIfBreaksDirection(direction, current, previous int) bool {
  difference := current - previous
  if direction < 0 {
    return difference > 0
  }
  return difference < 0
}

func main() {
  array := []int{-1,-5,-10,-1100,-1100,-1101,-1102,-9001}
  fmt.Printf("First result => %v\n", FirstSolution(array))
}
