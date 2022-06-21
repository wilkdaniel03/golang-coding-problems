package main

import "fmt"

func IterativeSolution(matrix [][]int) []int {
  result := []int{}
  firstRow, lastRow := 0, len(matrix)-1
  firstCol, lastCol := 0, len(matrix[0])-1
  for firstRow <= lastRow && firstCol <= lastCol {
    for col := firstCol; col <= lastCol; col++ {
      result = append(result, matrix[firstRow][col])
    }
    for row := firstRow+1; row <= lastRow; row++ {
      result = append(result, matrix[row][lastCol])
    }
    for col := lastCol-1; col >= firstCol; col-- {
      if firstRow == lastCol {
        break
      }
      result = append(result, matrix[lastRow][col])
    }
    for row := lastRow-1; row > firstRow; row-- {
      if firstCol == lastCol {
        break
      }
      result = append(result, matrix[row][firstCol])
    }
    firstCol++
    lastCol--
    firstRow++
    lastRow--
  }
  return result
}

func RecursiveSolution(matrix [][]int) []int {
  result := []int{}
  helper(0, len(matrix)-1, 0, len(matrix[0])-1, matrix, &result)
  return result
}

func helper(
  firstRow, lastRow, firstCol, lastCol int,
  matrix [][]int,
  result *[]int,
) {
  if firstRow > lastRow || firstCol > lastCol {
    return
  }
  for col := firstCol; col <= lastCol; col++ {
    *result = append(*result, matrix[firstRow][col])
  }
  for row := firstRow+1; row <= lastRow; row++ {
    *result = append(*result, matrix[row][lastCol])
  }
  for col := lastCol-1; col >= firstCol; col-- {
    if firstRow == lastRow {
      break
    }
    *result = append(*result, matrix[lastRow][col])
  }
  for row := lastRow-1; row > firstRow; row-- {
    if firstCol == lastCol {
      break
    }
    *result = append(*result, matrix[row][firstCol])
  }
  helper(
    firstRow+1,
    lastRow-1,
    firstCol+1,
    lastCol-1,
    matrix,
    result,
    )
}

func main() {
  matrix := [][]int{
    {1,2,3,4},
    {12,13,14,5},
    {11,16,15,6},
    {10,9,8,7},
  }
  fmt.Printf("First result => %v\n", IterativeSolution(matrix))
  fmt.Printf("Second result => %v\n", RecursiveSolution(matrix))
}
