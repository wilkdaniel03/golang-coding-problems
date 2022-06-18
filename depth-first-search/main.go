package main

import "fmt"

type BST struct {
	Value    string
	Children []*BST
}

func (tree *BST) Solution() []string {
  collection := []string{}
  solution(tree, &collection)
  return collection
}

func solution(node *BST, collection *[]string) {
  *collection = append(*collection, node.Value)
  for _, n := range node.Children {
    solution(n, collection)
  }
}

func main() {
  root := BST{Value: "A", Children: nil}
  midNodeOne := BST{Value: "B", Children: nil}
  midNodeTwo := BST{Value: "C", Children: nil}
  midNodeThree := BST{Value: "D", Children: nil}
  bottomNodeOne := BST{Value: "E", Children: nil}
  bottomNodeTwo := BST{Value: "F", Children: nil}
  bottomNodeThree := BST{Value: "G", Children: nil}
  bottomNodeFour := BST{Value: "H", Children: nil}
  root.Children = []*BST{&midNodeOne, &midNodeTwo, &midNodeThree}
  midNodeOne.Children = []*BST{&bottomNodeOne, &bottomNodeTwo}
  midNodeThree.Children = []*BST{&bottomNodeThree, &bottomNodeFour}
  fmt.Printf("Result => %v", root.Solution())
}
