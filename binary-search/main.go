package main

import "fmt"

type BST struct {
	Value       int
  Index       int
	Left, Right *BST
}

func (tree *BST) Insert(value, index int) {
  newNode := BST{Value: value, Index: index, Left: nil, Right: nil}
  currentNode := tree
  for currentNode != nil {
    if newNode.Value < currentNode.Value {
      if currentNode.Left == nil {
        currentNode.Left = &newNode
        break
      } else {
        currentNode = currentNode.Left
      }
    } else {
      if currentNode.Right == nil {
        currentNode.Right = &newNode
        break
      } else {
        currentNode = currentNode.Right
      }
    }
  }
}

func (tree *BST) Lookup(value int) int {
  currentNode := tree
  for currentNode != nil {
    if value < currentNode.Value {
      currentNode = currentNode.Left
    } else if value > currentNode.Value {
      currentNode = currentNode.Right
    } else {
      return currentNode.Index
    }
  }
  return -1
}

func main() {
  root := BST{Value: 0, Index: 0, Left: nil, Right: nil}
  root.Insert(1, 1)
  root.Insert(21, 2)
  root.Insert(33, 3)
  root.Insert(45, 4)
  root.Insert(46, 5)
  root.Insert(61, 6)
  root.Insert(71, 7)
  root.Insert(72, 8)
  root.Insert(73, 8)
  fmt.Printf("Result => %d\n", root.Lookup(33))

}
