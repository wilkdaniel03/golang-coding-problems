package main

import "fmt"

func WorstSolution(str string) bool {
  reversedString := ""
  for i := len(str)-1; i >= 0 ; i-- {
    reversedString += string(str[i])
  }
  return str == reversedString
}

func AverageIterativeSolution(str string) bool {
  charArray := []byte{}
  for i := len(str)-1; i >= 0; i-- {
    charArray = append(charArray, str[i])
  }
  return str == string(charArray)
}

func main() {
  fmt.Printf("First result => %v\n", WorstSolution("abcdcba"))
  fmt.Printf("Second result => %v\n", AverageIterativeSolution("abcdcba"))
}
