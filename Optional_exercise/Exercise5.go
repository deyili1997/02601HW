package main
import (
  "fmt"
  "math/rand"
)
func main() {
  fmt.Println(sample([]float64{0.2,0.1,0.3,0.4}))
}

func sample(dis []float64) int {
  posibilityList := make([]float64,len(dis)+1)
  posibilityList[0] = 0.0
  for i:= 1; i<=len(dis);i++ {
    posibilityList[i] = dis[i-1] + posibilityList[i-1]
  }
  fmt.Println(posibilityList)
  choice := rand.Float64()
  fmt.Println(choice)
  for m := 0; m<len(posibilityList)-1; m++ {
    if choice >= posibilityList[m] && choice < posibilityList[m+1] {
      return m
    }
  }
  panic("You should not go here")
  return -1
}
