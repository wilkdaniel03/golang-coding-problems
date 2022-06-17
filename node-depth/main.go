package main

import "fmt"

type BST struct {
	Value int
	Left  *BST
	Right *BST
}

type Level struct {
  Node *BST
  Depth int
}

func (tree *BST) IterativeSolution() int {
  depth := 0
  var top Level
  stack := []Level{{Node: tree, Depth: 0}}
  for len(stack) > 0 {
    top, stack = stack[len(stack)-1], stack[:len(stack)-1]
    node, currentDepth := top.Node, top.Depth
    if node == nil {
      continue
    }
    depth += currentDepth
    stack = append(stack, Level{Node: node.Left, Depth: currentDepth + 1})
    stack = append(stack, Level{Node: node.Right, Depth: currentDepth + 1})
  }
  return depth
}

func (tree *BST) RecursiveSolution() int {
  return recursiveSolution(tree, 0)
}

func recursiveSolution(node *BST, depth int) int {
  if node == nil {
    return 0
  }
  return depth + recursiveSolution(node.Left, depth+1) + recursiveSolution(node.Right, depth+1)
}

func main() {
  tree := BST{Value: 10, Left: nil, Right: nil}
  midTreeOne := BST{Value: 5, Left: nil, Right: nil}
  midTreeTwo := BST{Value: 15, Left: nil, Right: nil}
  tree.Left = &midTreeOne
  tree.Right = &midTreeTwo
  bottomTreeOne := BST{Value: 2, Left: nil, Right: nil}
  bottomTreeTwo := BST{Value: 5, Left: nil, Right: nil}
  bottomTreeThree := BST{Value: 13, Left: nil, Right: nil}
  bottomTreeFour := BST{Value: 22, Left: nil, Right: nil}
  midTreeOne.Left = &bottomTreeOne
  midTreeOne.Right = &bottomTreeTwo
  midTreeTwo.Left = &bottomTreeThree
  midTreeTwo.Right = &bottomTreeFour
  fmt.Printf("Iterative => %d\n", tree.IterativeSolution())
  fmt.Printf("Recursive => %d\n", tree.RecursiveSolution())
}
