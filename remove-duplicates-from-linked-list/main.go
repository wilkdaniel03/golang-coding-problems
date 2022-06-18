package main

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func (ll *LinkedList) Solution() LinkedList {
  currentNode := ll
  for currentNode != nil {
    nextDistinct := currentNode.Next
    for nextDistinct != nil && nextDistinct.Value == currentNode.Value {
      nextDistinct = nextDistinct.Next
    }
    currentNode.Next = nextDistinct
    currentNode = nextDistinct
  }
  return *currentNode
}

func main() {

}
