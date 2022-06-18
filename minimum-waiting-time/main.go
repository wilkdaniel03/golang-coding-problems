package main

import (
	"fmt"
	"sort"
)

func Solution(queries []int) int {
  sort.Ints(queries)
  currentDuration := 0
  for i, q := range queries {
    queriesAfterwards := len(queries) - (i+1)
    currentDuration += q * queriesAfterwards
  }
  return currentDuration
}

func main() {
  sampleQueries := []int{3,2,1,2,6}
  fmt.Printf("Result => %d\n", Solution(sampleQueries))
}
