package main
import (
  "fmt"
)

func main() {
  // // Question1
  // T1 := []bool{true, true, false, false, true, false}
  // T2 := []bool{}
  // fmt.Println(CountSwitches(T1))
  // fmt.Println(CountSwitches(T2))



  // Question2
  // T2 := [][]int{
  //         []int{0,4,-5,1,7},
  //         []int{-10,10,10},
  //         []int{0, 1, 0, -1},
  //         []int{3,-3,-3,0},
  //         []int{4,-4,5,0},
  //     }
  // fmt.Println(LeftmostNegative(T2))


  // //Question3
  // T3 := []int{7, 8, 2, 8, 2, 7, 8}
  // fmt.Println(FirstUnique(T3))


  // //Quesion4
  // T4A := []int{2,2,2,3}
  // T4B := []int{3,2,2,2}
  // fmt.Println(IsCircularPermutation(T4A,T4B))


  // //Question5
  // T5A := []int{1,23,7,9,7,7,7,5,5,5,4,4,4,1,1}
  // T5B := []int{1,7,8, 8, 7,4,5}
  // fmt.Println(Contains(T5A,T5B))


  // //Quesion6
  // T5 := []int{1,5,9,9,20,1,21,3,4}
  // fmt.Println(SquareNumbers(T5))


  // // Question7
  // var n, n1, n2, n3, n4, n5, n6, n7 Node
  // n.next = &n1
  // n1.next = &n2
  // n2.next = &n3
  // n3.next = &n4
  // n4.next = &n5
  // n5.next = &n6
  // n6.next = &n7
  // n7.next = nil
  // fmt.Println(HasCycle(&n))


  fmt.Println("=======================第一题===================================")
  booleanTest1 := []bool{true, true, false, false, true, false}
  fmt.Println(CountSwitches(booleanTest1))
  booleanTest2 := []bool{true, true, false, false, true, false, true, false}
  fmt.Println(CountSwitches(booleanTest2))
  fmt.Println(CountSwitches([]bool{}))



  fmt.Println("=======================第二题===================================")
	test1 := [][]int{
		{3, 4, -5, 1, 7},
		{-10, 10, 10},
		{2, 1, 0, -1},
		{3, -3, -3, 0},
		{4, -4, 5, 0},
	}
	fmt.Println(LeftmostNegative(test1))

	test2 := [][]int{
		{3, 4, 5, 1, 7},
		{10, 10, 10},
		{2, 1, 0, 1},
		{3, 3, 3, 0},
		{4, 4, 5, 0},
	}
	fmt.Println(LeftmostNegative(test2))

  fmt.Println(LeftmostNegative([][]int{}))

  fmt.Println("=======================第三题===================================")
	fmt.Println(FirstUnique([]int{2, 3, 4, 5, 2, 4, 5}))
	fmt.Println(FirstUnique([]int{7, 8, 2, 8, 2, 7, 8}))
	fmt.Println(FirstUnique([]int{8, 8, 8, 9, 8, 6, 8, 8}))
  fmt.Println(FirstUnique([]int{}))

  fmt.Println("=======================第四题===================================")
	// true
	a := []int{2, 4, 5, 6, 7, 2}
	b := []int{4, 5, 6, 7, 2, 2}
	//b := []int{7, 8, 1, 10, 14, 31} // false case

	//true
	c := []int{2, 3, 3, 4}
	// d := []int{3, 3, 4, 2}
	//d := []int{3, 4, 4, 2}

	fmt.Println(IsCircularPermutation(a, b))
	fmt.Println(IsCircularPermutation(c, []int{}))
	fmt.Println(IsCircularPermutation([]int{2, 2, 2, 3}, []int{3, 2, 2, 2}))

  fmt.Println("=======================第五题===================================")
  // l2's 10 occurs 3 times, while l1 10 occurs 2 times => false
  l1 := []int{1, 7, 8, 10, 10, 31, 14}
  l2 := []int{10, 10, 10, 31, 14, 1, 7, 8}
  fmt.Println(Contains(l1, l2))

  // true
  l3 := []int{1, 7, 8, 10, 31, 14}
  l4 := []int{10, 31, 14, 1, 7, 8}
  fmt.Println(Contains(l3, l4))

  // true
  l5 := []int{0, 1, 7, 8, 10, 10, 31, 14}
  l6 := []int{10, 31, 14, 1, 7, 8}
  fmt.Println(Contains(l5, l6))

  fmt.Println(Contains([]int{},l6))
  fmt.Println(Contains([]int{},[]int{}))
  fmt.Println(Contains(l6,[]int{}))


  fmt.Println("=======================第六题===================================")
  nums := []int{0, 1, 5, 9, 9, 20}
	fmt.Println(SquareNumbers(nums))
	nums2 := []int{0, 1, 5, 9, 9, 20, 21, 23, 36}
	fmt.Println(SquareNumbers(nums2))
	nums3 := []int{1, 5, 9, 9, 20,453463454523424}
	fmt.Println(SquareNumbers(nums3))

  fmt.Println(SquareNumbers([]int{}))

}
