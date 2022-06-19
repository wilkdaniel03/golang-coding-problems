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

func IterativeSolution(n int) int {
  currentFib := [2]int{0,1}
  for i := 3; i <= n; i++ {
    nextFib := currentFib[0] + currentFib[1]
    currentFib = [2]int{currentFib[1], nextFib}
  }
  if n == 0 {
    return currentFib[0]
  }
  return currentFib[1]
}

func main() {
	fmt.Printf("Factorial Result => %d\n", FactorialSolution(20))
	fmt.Printf("Cached Result => %d\n", CachedSolution(20))
	fmt.Printf("Iterative Result => %d\n", IterativeSolution(20))
}
