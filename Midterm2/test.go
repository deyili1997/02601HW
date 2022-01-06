package main

import "fmt"

func main() {
	//****************Q1**Q2*******************************
	// var N0, N1, N2, N3, N4, N5, N6, N7, N8, N9 TreeNode
	// N0.value = 0
	// N1.value = 1
	// N2.value = 2
	// N3.value = 3
	// N4.value = 4
	// N5.value = 5
	// N6.value = 6
	// N7.value = 7
	// N8.value = 8
	// N9.value = 9
	// N5.left = &N3
	// N5.right = &N7
	// N3.left = &N1
	// N3.right = &N4
	// N1.left = &N0
	// N1.right = &N2
	// N7.left = &N6
	// N7.right = &N8
	// N8.right = &N9
	// fmt.Println(IsBSTOrdered(&N5))
	// fmt.Println(FindPath(&N5, 9))
	// fmt.Println(FindPath(&N5, 6))
	// fmt.Println(LowestCommonAncestor(&N5, 6, 9))
	// *********************Q1Q2**********************//
	// var N2, N3, N4, N5, N6 TreeNode
	// N2.value = 2
	// N2.left = nil
	// N2.right = &N3
	// N3.value = 3
	// N3.left = nil
	// N3.right = nil
	// N4.value = 4
	// N4.left = &N2
	// N4.right = &N5
	// N5.value = 5
	// N5.left = nil
	// N5.right = &N6
	// N6.value = 6
	// N6.left = nil
	// N6.right = nil
	// fmt.Println(FindChildren(&N4))
	// fmt.Println(IsBSTOrdered(&N4))
	// fmt.Println(LowestCommonAncestor(&N4, 5, 6))
	//**********************Q1Q2******************//
	// var N1, N2 TreeNode
	// N2.value = 2
	// N2.left = nil
	// N2.right = nil
	// N1.value = 1
	// N1.left = nil
	// N1.right = &N2
	// fmt.Println(IsBSTOrdered(&N1))
	// fmt.Println("Finsh IsBSTOrdered")
	// fmt.Println(FindPath(&N1, 0))
	// fmt.Println()
	// fmt.Println(FindPath(&N1, 1))
	// fmt.Println()
	// fmt.Println(FindPath(&N1, 2))
	// fmt.Println()
	// fmt.Println(LowestCommonAncestor(&N1, 2, 1))
	// fmt.Println("Finsh LowestCommonAncestor")
	//********************Q1Q2*************
	// var N2, N3, N4 TreeNode
	// N4.value = 4
	// N4.left = &N3
	// N4.right = nil
	// N3.value = 3
	// N3.left = &N2
	// N3.right = nil
	// N2.value = 2
	// N2.left = nil
	// N2.right = nil
	// fmt.Println(IsBSTOrdered(&N4))
	// fmt.Println(LowestCommonAncestor(&N4, 2, 4))
	//*****************Q3******************//
	// var l1, l2, l3, l4, l5, l6, l7, l8, l9, l10, l11, l12, l13, l14, l15, l16, l17, l18 ListNode
	// l1.value = 1
	// l1.next = &l4
	// l2.value = 2
	// l2.next = &l3
	// l3.value = 3
	// l3.next = &l6
	// l4.value = 4
	// l4.next = &l5
	// l5.value = 5
	// l5.next = &l10
	// l6.value = 6
	// l6.next = &l7
	// l7.value = 7
	// l7.next = &l8
	// l8.value = 8
	// l8.next = &l9
	// l9.value = 9
	// l9.next = &l12
	// l10.value = 10
	// l10.next = &l11
	// l11.value = 11
	// l11.next = &l13
	// l12.value = 12
	// l12.next = &l15
	// l13.value = 13
	// l13.next = &l14
	// l14.value = 14
	// l14.next = &l16
	// l15.value = 15
	// l15.next = nil
	// l16.value = 16
	// l16.next = &l17
	// l17.value = 17
	// l17.next = &l18
	// l18.value = 18
	// l18.next = nil
	// result := MergeSortedLists(&l2, &l1)
	// fmt.Println(result.value)
	// fmt.Println(result.next.value)
	// fmt.Println(result.next.next.value)
	// fmt.Println(result.next.next.next.value)
	// fmt.Println(result.next.next.next.next.value)
	// fmt.Println(result.next.next.next.next.next.value)
	// fmt.Println(result.next.next.next.next.next.next.value)
	// fmt.Println(l8.next.value)
	// fmt.Println(l17.next.value)
	// fmt.Println(l12.next.value)
	// fmt.Println(l3.next.value)
	// ************************Q3***********************************
	// var l1, l2, l3, l4, l8 ListNode
	// l1.value = 1
	// l1.next = &l8
	// l2.value = 2
	// l2.next = &l3
	// l3.value = 3
	// l3.next = &l4
	// l4.value = 4
	// l4.next = nil
	// l8.value = 8
	// l8.next = nil
	// a := MergeSortedLists(&l1, &l2)
	// fmt.Println(a.value)
	// fmt.Println(a.next.value)
	// fmt.Println(a.next.next.value)
	// fmt.Println(a.next.next.next.value)
	// fmt.Println(a.next.next.next.next.value)

	//********************Q3**************************************
	// var l1, l2, l3, l4, l5, l6 ListNode
	// l1.value = 1
	// l1.next = &l2
	// l2.value = 3
	// l2.next = &l3
	// l3.value = 4
	// l3.next = nil
	// l4.value = 1
	// l4.next = &l5
	// l5.value = 2
	// l5.next = &l6
	// l6.value = 4
	// l6.next = nil
	// MergeSortedLists(&l1, &l4)
	// fmt.Println(l1.value)
	// fmt.Println(l1.next.value)
	// fmt.Println(l1.next.next.value)
	// fmt.Println(l1.next.next.next.value)
	// fmt.Println(l1.next.next.next.next.value)
	// fmt.Println(l1.next.next.next.next.next.value)

	//**************************Q3*********************************
	// var l1, l2, l3 ListNode
	// var n *ListNode
	// l1.value = 1
	// l1.next = &l2
	// l2.value = 2
	// l2.next = &l3
	// l3.value = 3
	// l3.next = nil
	// n = nil
	// a := MergeSortedLists(n, &l1)
	// fmt.Println(a.next.next.value)
	//*************************Q3********************************
	// var l1, l2, l3, l4 ListNode
	// l1.value = 1
	// l1.next = &l2
	// l2.value = 2
	// l2.next = &l3
	// l3.value = 3
	// l3.next = nil
	// l4.value = 4
	// l4.next = nil
	// a := MergeSortedLists(&l4, &l1)
	// fmt.Println(a.value)
	// fmt.Println(a.next.value)
	// fmt.Println(a.next.next.value)
	// fmt.Println(a.next.next.next.value)

	//***************************Q3********************************
	// var n1, n2 *ListNode
	// n1 = nil
	// n2 = nil
	// fmt.Println(MergeSortedLists(n1, n2))
	//***********************Q3*******************************
	// var n2 ListNode
	// var n1 *ListNode
	// n2.value = 1
	// n2.next = nil
	// n1 = nil
	// fmt.Println(MergeSortedLists(n1, &n2).value)
	//*************************Q4*******************************
	// var l1, l2, l3, l4, l5, l6 ListNode
	// l1.value = 8
	// l1.next = &l2
	// l2.value = 4
	// l2.next = &l3
	// l3.value = 3
	// l3.next = &l4
	// l4.value = 0
	// l4.next = &l5
	// l5.value = 5
	// l5.next = &l6
	// l6.value = 7
	// l6.next = nil
	// fmt.Println(MthFromEnd(&l1, 5))
	// **********************Q4*****************************
	// var l1 ListNode
	// l1.value = 0
	// l1.next = nil
	// fmt.Println(MthFromEnd(&l1, 0))
	//****************************Q5*********************
	// var S = Stack{}
	// S.Push(5)
	// S.Push(2)
	// S.Push(3)
	// S.Push(1)
	// S.Push(8)
	// fmt.Println(S.Min()) //1
	// pop1 := S.Pop()
	// fmt.Println("The pop1 is ", pop1)
	// fmt.Println(S.Min()) //1
	// pop2 := S.Pop()
	// fmt.Println("The pop2 is", pop2)
	// fmt.Println(S.Min()) //2
	// pop3 := S.Pop()
	// fmt.Println("The pop3 is", pop3)
	// fmt.Println(S.Min()) //2
	// pop4 := S.Pop()
	// fmt.Println("The pop4 is", pop4)
	// fmt.Println(S.Min()) //5
	// pop5 := S.Pop()
	// fmt.Println("The pop5 is", pop5)
	//*******************Q6***********************
	r1 := []int{1, 2, 3, 4, 5, 6}
	r2 := []int{2, 3, 4, 5, 6, 7}
	r3 := []int{4, 2, 3, 6, 7, 8}
	r4 := []int{1, 2, 3, 4, 1, 9}
	r5 := []int{1, 2, 8, 1, 1, 2}
	test := [][]int{r1, r2, r3, r4, r5}
	for r := range test {
		fmt.Println(test[r])
	}
	delete := RemoveRowsColumns(test, 1)
	fmt.Println("Now finish deleting!")
	for r := range delete {
		fmt.Println(delete[r])
	}
	//***************Q6*****************************
	// 	r1 := []int{1, 2}
	// 	r2 := []int{2, 1}
	// 	test := [][]int{r1, r2}
	// 	for r := range test {
	// 		fmt.Println(test[r])
	// 	}
	// 	delete := RemoveRowsColumns(test, 1)
	// 	fmt.Println("Now finish deleting!")
	// 	fmt.Println(delete)
	//******************Q6*************************
	// r1 := []int{1, 2}
	// test := [][]int{r1}
	// delete := RemoveRowsColumns(test, 2)
	// fmt.Println(delete)
}
