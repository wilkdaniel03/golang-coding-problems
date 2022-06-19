package main

import (
	"fmt"
	"strconv"
)

func Solution(str string) string {
  bytes := []byte{}
  currentLength := 1
  for i := 1; i < len(str); i++ {
    currentChar, previousChar := str[i], str[i-1]
    if currentChar != previousChar || currentLength == 9 {
      bytes = append(bytes, strconv.Itoa(currentLength)[0])
      bytes = append(bytes, currentChar)
      currentLength = 0
    } 
    currentLength += 1
  }
  bytes = append(bytes, strconv.Itoa(currentLength)[0])
  bytes = append(bytes, str[len(str)-1])
  return string(bytes)
}

func main() {
  string := "AAAAAAAAAAAAAAAAAAAAAAAAABBBBCC"
  fmt.Printf("Result => %s\n", Solution(string))
}
