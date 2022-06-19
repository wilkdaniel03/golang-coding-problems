package main

import "fmt"

func FactorialSolution(n int) int {
	if n == 1 {
		return 0
	} else if n == 2 {
		return 1
	} else {
		return FactorialSolution(n-1) + FactorialSolution(n-2)
	}
}

func CachedSolution(n int) int {
	cache := map[int]int{1: 0, 2: 1}
  return findFib(n, cache)
}

func findFib(n int, cache map[int]int) int {
  if val, found := cache[n]; found {
    return val
  }
  return findFib(n-1, cache) + findFib(n-2, cache)
}

func main() {
	fmt.Printf("Factorial Result => %d\n", FactorialSolution(20))
	fmt.Printf("Cached Result => %d\n", CachedSolution(20))
}
