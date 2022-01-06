package main
import (
  "fmt"
)

func main(){
  a := [][]bool{
    []bool{true,false,true,true},
    []bool{true,false,true,false},
    []bool{false,true,false,false},
    []bool{true,false,false,true},
  }
  fmt.Println(connectFour(a))
}
func connectFour(board [][]bool) bool {
  length := len(board)
  for r := range board {
    for c := range board[r]{
//check rows
      if c + 3 <= length -1{
        if board[r][c] == true && board[r][c+1] == true && board[r][c+2] == true && board[r][c+3] == true {
          return true
        }
      }

//check columns
      if r + 3 <= length -1{
        if board[r][c] == true && board[r+1][c] == true && board[r+2][c] == true && board[r+3][c] == true {
          return true
        }
      }
//check diagonal 1
      if r == c && r + 3 < length && c + 3 < length{
        if board[r][c] == true && board[r+1][c+1] == true && board[r+2][c+2] == true && board[r+3][c+3] == true {
          return true
      }
    }
//check diagonal 2
      if c == length - r - 1 && c - 3 >= 0 && r + 3 < length {
        if board[r][c] == true && board[r+1][c-1] == true && board[r+2][c-2] == true && board[r+3][c-3]== true {
          return true
          }
        }
      }
    }
    return false
}
