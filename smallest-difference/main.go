package main

import (
	"fmt"
	"math"
	"sort"
)

func Solution(arrayOne, arrayTwo []int) []int {
  sort.Ints(arrayOne)
  sort.Ints(arrayTwo)
  smallest, current, smallestPair := math.MaxInt32, math.MaxInt32, []int{}
  idxOne, idxTwo := 0, 0
  for idxOne < len(arrayOne) && idxTwo < len(arrayTwo) {
    firstNum, secondNum := arrayOne[idxOne], arrayTwo[idxTwo]
    current = abs(firstNum - secondNum)
    if firstNum < secondNum {
      idxOne += 1
    } else if secondNum < firstNum {
      idxTwo += 1
    } else {
      return []int{firstNum, secondNum}
    }
    if smallest > current {
      smallest = current
      smallestPair = []int{firstNum, secondNum}
    }
  }
  return smallestPair
}

func abs(a int) int {
  if a < 0 {
    return -a
  }
  return a
}

func main() {
  arrayOne := []int{-1,5,10,20,28,3}
  arrayTwo := []int{26,134,135,15,17}
  fmt.Printf("Result => %v\n", Solution(arrayOne, arrayTwo))
}
