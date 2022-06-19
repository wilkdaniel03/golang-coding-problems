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

func OptimalSolution(str string) int {
  chars := map[rune]int{}
  for _, char := range str {
    if _, ok := chars[char]; !ok {
      chars[char] = 1
    } else {
      chars[char] += 1
    }
  }
  for idx, char := range str {
    if chars[char] == 1 {
      return idx
    }
  }
  return -1
}

func main() {
  string := "abcdcaf"
  fmt.Printf("First result => %d\n", NSquaredSolution(string))
  fmt.Printf("Second result => %d\n", OptimalSolution(string))
}
