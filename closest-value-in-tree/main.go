package main

import "fmt"

type BST struct {
	Value int
	Left  *BST
	Right *BST
}

func (tree *BST) IterativeSolution(target int) int {
  return tree.iterativeSolution(target, tree.Value)
}

func (tree *BST) iterativeSolution(target, closest int) int {
  currentNode := tree
  for currentNode != nil {
    if abs(target - closest) > abs(target - currentNode.Value) {
      closest = currentNode.Value
    }
    if target < currentNode.Value {
      currentNode = currentNode.Left
    } else if target > currentNode.Value {
      currentNode = currentNode.Right
    } else {
      return closest
    }
  }
  return closest
}

func (tree *BST) RecursiveSolution(target int) int {
  return tree.recursiveSolution(target, tree.Value)
}

func (tree *BST) recursiveSolution(target, closest int) int {
  if abs(target - closest) > abs(target - tree.Value) {
    closest = tree.Value
  }
  if target < tree.Value && tree.Left != nil {
    return tree.Left.recursiveSolution(target, closest)
  } else if target > tree.Value && tree.Right != nil {
    return tree.Right.recursiveSolution(target, closest)
  }
  return closest
}

func abs(a int) int {
  if a < 0 {
    return -a
  }
  return a
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
  fmt.Printf("Iterative solution => %d\n", tree.IterativeSolution(12))
  fmt.Printf("Recursive solution => %d\n", tree.RecursiveSolution(12))
}
