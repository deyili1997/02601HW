package main
import (
  "fmt"
)

func main() {
  a := []int{2,3,4}
  // fmt.Println(Sum(a, 4))
  wayCount = 0
  countPartitions(a, 6)
  fmt.Println(wayCount)
}


func countPartitions(L []int, k int) int {
  if k == 0 {
    return 1
  }
  if k < 0 {
    return -1
  }
  for k,j := range L {
    g=L[k:]
    countPartitions(g, k-j)
    }
  return 1
}

//
// func Sum (L []int, n int) int {
//   if n == 1 {
//     return L[0]
//   }
//   if n > 1 {
//     return Sum(L, n-1) + L[n-1]
//   }
//   return 1
// }
