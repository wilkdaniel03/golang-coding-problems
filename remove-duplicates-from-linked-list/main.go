package main

import "fmt"

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func (ll *LinkedList) Solution() *LinkedList {
	currentNode := ll
	for currentNode != nil {
		nextDistinct := currentNode.Next
		for nextDistinct != nil && nextDistinct.Value == currentNode.Value {
			nextDistinct = nextDistinct.Next
		}
		currentNode.Next = nextDistinct
		currentNode = nextDistinct
	}
  return currentNode
}

func (ll *LinkedList) Iterate() []int {
  collection := []int{}
  currentNode := ll
  for currentNode != nil {
    collection = append(collection, currentNode.Value)
    currentNode = currentNode.Next
  }
  return collection
}

func main() {
	head := LinkedList{Value: 1, Next: nil}
	secondNode := LinkedList{Value: 1, Next: nil}
	thirdNode := LinkedList{Value: 3, Next: nil}
	fourthNode := LinkedList{Value: 4, Next: nil}
	fifthNode := LinkedList{Value: 4, Next: nil}
	sixthNode := LinkedList{Value: 4, Next: nil}
	seventhNode := LinkedList{Value: 5, Next: nil}
	eigthNode := LinkedList{Value: 6, Next: nil}
	ninethNode := LinkedList{Value: 6, Next: nil}

  head.Next = &secondNode
  secondNode.Next = &thirdNode
  thirdNode.Next = &fourthNode
  fourthNode.Next = &fifthNode
  fifthNode.Next = &sixthNode
  sixthNode.Next = &seventhNode
  seventhNode.Next = &eigthNode
  eigthNode.Next = &ninethNode

  head.Solution()

  fmt.Printf("Result => %v\n", head.Iterate())
}
