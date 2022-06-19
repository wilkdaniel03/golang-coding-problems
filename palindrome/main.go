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

func AverageRecursiveSolution(str string) bool {
  return reversiveHelper(str, 0)
}

func reversiveHelper(str string, i int) bool {
  if i >= len(str)-1-i {
    return true
  }
  if str[i] != str[len(str)-1-i] {
    return false
  }
  return reversiveHelper(str, i+1)
}

func OptimalSolution(str string) bool {
  for startIdx := 0; startIdx < len(str); startIdx++ {
    if str[startIdx] != str[len(str)-1-startIdx] {
      return false
    }
  }
  return true
}

func main() {
  fmt.Printf("First result => %v\n", WorstSolution("abcdcba"))
  fmt.Printf("Second result => %v\n", AverageIterativeSolution("abcdcba"))
  fmt.Printf("Third result => %v\n", AverageRecursiveSolution("abcdcba"))
  fmt.Printf("Fourth result => %v\n", OptimalSolution("abcdcba"))
}
