package main

import (
	"fmt"
	"strings"
)

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

func WithoutUnicodesSolution(str string, key int) string {
  runes := []rune(str)
  alphabet := "abcdefghijklmnopqrstuvwxyz"
  for i, char := range(runes) {
    index := strings.Index(alphabet, string(char))
    if index == -1 {
      return ""
    }
    newIndex := (index+key) % 26
    runes[i] = rune(alphabet[newIndex])
  }
  return string(runes)
}

func main() {
  fmt.Printf("First result => %s\n", WithUnicodesSolution("xyz", 2))
  fmt.Printf("Second result => %s\n", WithoutUnicodesSolution("xyz", 2))
}
