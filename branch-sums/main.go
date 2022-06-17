package main

import "fmt"

type BST struct {
	Value int
	Left  *BST
	Right *BST
}

func (tree *BST) Solution() []int {
  sums := []int{}
  tree.solution(0, &sums)
  return sums
}

func (tree *BST) solution(currentSum int, sums *[]int) {
  if tree == nil {
    return
  }

  currentSum += tree.Value
  if tree.Left == nil && tree.Right == nil {
    *sums = append(*sums, currentSum)
    return
  }

  tree.Left.solution(currentSum, sums)
  tree.Right.solution(currentSum, sums)
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
  fmt.Printf("Solution => %d\n", tree.Solution())
}
