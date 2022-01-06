package main
import (
  "fmt"
)

func main() {
  a := []int{1,1,1,1,4,14,3,7,5}
  fmt.Println(secondSmallest(a))
}

func secondSmallest(L []int)int {
  bubbleSort(L)
  return find2(L)
}

func bubbleSort(L []int)string {
  for i := 0; i < len(L); i++ {
    for j := i+1; j< len(L) - i; j++ {
      if L[j] < L[j-1] {
        L[j],L[j-1] = L[j-1],L[j]
      }
    }
  }
  return "finished sort!"
}

func find2 (L []int) int {
  for i := 1; i < len(L); i++ {
    if L[i] != L[0] {
      return L[i]
    }
  }
  panic("You should not go here!")
  return 1
}
