package main
import (
  "fmt"
)
func main() {
  a := []int{1,4,7}
  b := []int{2,3,5,6,7,8,9}
  fmt.Println(Merge(a,b))
  fmt.Println(Merge2(a,b))
}
//Solution 1: Bubble sort
func Merge(L1, L2 []int) []int {
  for _,i := range L2 {
    L1 = append(L1, i)
  }
  for m := 0; m < len(L1); m++ {
    for n := 1; n < len(L1) - m; n++ {
      if L1[n] < L1[n-1] {
        L1[n],L1[n-1] = L1[n-1], L1[n]
      }
    }
  }
  return L1
}
// Solution 2: two pointers
func Merge2(L1, L2 []int) []int {
  mergedList := make([]int,0)
  pointer1 := 0
  pointer2 := 0
  for len(mergedList) < len(L1) + len(L2) {
    if pointer1 == len(L1) {
      fmt.Println("L1 has finished .Now append the rest of L2, and the pointer2 is now",pointer2)
      for i := pointer2; i< len(L2); i++{
        mergedList = append(mergedList,L2[i])
      }
      break
    }
    if pointer2 == len(L2) {
      fmt.Println("L2 has finished .Now append the rest of L1, and the pointer1 is now", pointer1)
      for j := pointer1; j< len(L1); j++{
        mergedList = append(mergedList,L1[j])
      }
      break
    }

    if L1[pointer1] <= L2[pointer2] {
      fmt.Println("L1[",pointer1,"]<L2[",pointer2,"],now append:",L1[pointer1])
      mergedList = append(mergedList, L1[pointer1])
      if pointer1 < len(L1) {
        pointer1++
      }
    } else {
      fmt.Println("L2[",pointer2,"]<L1[",pointer1,"],now append:",L2[pointer2])
      mergedList = append(mergedList, L2[pointer2])
      if pointer2 < len(L2) {
        pointer2++
      }
    }
  }
  return mergedList
}
