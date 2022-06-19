package main

import "fmt"

func NSquaredSolution(str string) int {
  for i := range str {
    isDuplicated := false 
    for j := range str {
      if i != j && str[i] == str[j] {
        isDuplicated = true
      }
    }
    if !isDuplicated {
      return i
    }
  }
  return -1
}

func main() {
  string := "abcdcaf"
  fmt.Printf("First result => %d\n", NSquaredSolution(string))
}
