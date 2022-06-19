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

func main() {
	fmt.Printf("Factorial Result => %d\n", FactorialSolution(20))
}
