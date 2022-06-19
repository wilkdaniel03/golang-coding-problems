package main

import "fmt"

func NSquaredSolution(characters string, document string) bool {
  for _, c := range document {
    expectedFreq := checkCharacter(c, document)
    freq := checkCharacter(c, characters)
    if expectedFreq > freq {
      return false
    }
  }
  return true
}

func checkCharacter(char rune, collection string) int {
  freq := 0
  for _, c := range collection {
    if c == char {
      freq += 1
    }
  }
  return freq
}

func main() {
  characters := "Bste!hetsi ogEAxpelrt x "
  document := "AlgoExpert is the Best!"
  fmt.Printf("First result => %v\n", NSquaredSolution(characters, document))
}
