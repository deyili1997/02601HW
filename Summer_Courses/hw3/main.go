package main
import (
  "fmt"
)

// Exercise 1: Write a function InitializeMatrix that takes two integer parameters numRows and numCols and returns a 2-dimensional matrix of integers with numRows rows and numCols columns, all of whose values are zero.
func main () {
  fmt.Println(InitializeMatrix(3,4))
  row1 := []float64 {0.1,0.2,0.3}
  row2 := []float64 {0.4,0.5,0.6}
  var trial [][]float64
  trial = append(trial,row1)
  trial = append(trial,row2)
  fmt.Println(BinarizeMatrix(trial, 0.5))
  row3 := []bool {true,true,false,true}
  row4 := []bool {true,false,true,true}
  row5 := []bool {false,true,false,false}
  row6 := []bool {true,true,false,false}
  var trial2 [][]bool
  trial2 = append(trial2,row3)
  trial2 = append(trial2,row4)
  trial2 = append(trial2,row5)
  trial2 = append(trial2,row6)
  fmt.Println(ConnectFour(trial2))
}

func InitializeMatrix (numRows, numCols int) [][]int {
  var board [][]int
  for r := 0; r < numRows; r++ {
    thisRow := make([]int, numCols)
    board = append(board, thisRow)
  }
  return board
}

// Exercise 2: Write a function BinarizeMatrix that takes a two-dimensional matrix mtx of decimal variables along with a parameter threshold. It should return a two-dimensional matrix B of Boolean variables having the same dimensions as mtx and such that a value of B is true if the corresponding value of mtx is at least equal to threshold.

func BinarizeMatrix(mtx [][]float64, threshold float64) [][]bool {
  var result [][]bool
  rowLength := len(mtx)
  colLength := len(mtx[0])
  result = InitializeMatrixBool(rowLength,colLength)
  for r := 0; r < rowLength; r++ {
    for c := 0; c < colLength; c++ {
      if mtx[r][c] >= threshold {
        result[r][c] = true
      }
    }
  }
  return result
}

func InitializeMatrixBool (numRows, numCols int) [][]bool {
  var board [][]bool
  for r := 0; r < numRows; r++ {
    thisRow := make([]bool, numCols)
    board = append(board, thisRow)
  }
  return board
}

// Exercise 3: Write a function ConnectFour that takes a two-dimensional matrix B of Boolean variables as input. It should return true if there are four consecutive values in the same row, column, or diagonal that have the same value; it should return false otherwise.

func ConnectFour(B [][]bool) bool {
  rowResult := CheckRow(B)
  colResult := CheckCol(B)
  diagResult := CheckDiag(B)
  if rowResult || colResult || diagResult {
    return true
    }
  return false
}

func CheckRow (B [][]bool) bool {
  var status bool
  for r:=0;r<len(B);r++ {
    status = true
    for c:=1;c<len(B[0]);c++ {
      if B[r][c] != B[r][c-1] {
        status = false
        break
      }
    }
    if status == true {
      return status
    }
  }
  return status
}

func CheckCol(B [][]bool) bool {
  var status bool
  for c:=0;c<len(B[0]);c++ {
    status = true
    for r:=1;r<len(B);r++ {
      if B[r][c] != B[r-1][c] {
        status = false
        break
      }
    }
    if status == true {
      return status
    }
  }
  return status
}

func CheckDiag(B [][]bool) bool {
  var status bool
  status = true
  for i:= 1; i < len(B[0]);i++ {
    if B[i][i] != B[i-1][i-1] {
      status = false
      break
  }
  }
  if status == true {
    return status
    } else {
      status = true
    }
  for j:= len(B[0])-2; j >= 0; j-- {
    if B[len(B)-j-1][j]!=B[len(B)-j-2][j+1] {
      status = false
      break
    }
  if status == true {
    return status
    }
  }
  return status
}
