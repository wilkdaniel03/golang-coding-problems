package main

import "fmt"

func WorstSolution(str string) bool {
  reversedString := ""
  for i := len(str)-1; i >= 0 ; i-- {
    reversedString += string(str[i])
  }
  return str == reversedString
}

func main() {
  fmt.Printf("First result => %v\n", WorstSolution("abcdcba"))
}
