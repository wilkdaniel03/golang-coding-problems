package main

import "fmt"

func WithUnicodesSolution(str string, key int) string {
  shift, offset := rune(key%26), rune(26)
  runes := []rune(str)
  for i, char := range(runes) {
    if char >= 'a' &&  char+shift <= 'z' {
      char += shift
    } else {
      char += shift - offset
    }
    runes[i] = char
  }
  return string(runes)
}

func main() {
  fmt.Printf("First result => %s\n", WithUnicodesSolution("xyz", 2))
}
