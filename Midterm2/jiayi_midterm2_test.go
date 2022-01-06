package main

import (
	"testing"
)

func Test_IsBSTOrdered(t *testing.T) {
	test := func(t testing.TB, got, want bool) {
		t.Helper()
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	}

	t.Run("BST is ordered", func(t *testing.T) {
		root := TreeNode{value:6}
		node1 := TreeNode{value: 3}
		node2 := TreeNode{value: 1}
		node3 := TreeNode{value: 8}
		node4 := TreeNode{value: 7}
		node5 := TreeNode{value: 15}

		root.left = &node1; root.right = &node3
		node1.left = &node2
		node3.left = &node4; node3.right = &node5

		got := IsBSTOrdered(&root)
		want := true
		test(t, got, want)
	})

	t.Run("BST is not ordered 1", func(t *testing.T) {
		root := TreeNode{value:6}
		node1 := TreeNode{value: 3}
		node2 := TreeNode{value: 1}
		node3 := TreeNode{value: 8}
		node4 := TreeNode{value: 15}
		node5 := TreeNode{value: 7}

		root.left = &node1; root.right = &node3
		node1.left = &node2
		node3.left = &node4; node3.right = &node5

		got := IsBSTOrdered(&root)
		want := false
		test(t, got, want)
	})

	t.Run("BST is not ordered 2", func(t *testing.T){
		root := TreeNode{value: 5}
		node1 := TreeNode{value: 3}
		node2 := TreeNode{value: 1}
		node3 := TreeNode{value: 4}
		node4 := TreeNode{value: 6}
		node5 := TreeNode{value: 7}
		node6 := TreeNode{value: 0}
		node7 := TreeNode{value: 2}
		node8 := TreeNode{value: 9}
		node9 := TreeNode{value: 8}

		root.left = &node1; root.right = &node5
		node1.left = &node2; node1.right = &node3
		node2.left = &node6; node2.right = &node7
		node4.right = &node8
		node5.left = &node4; node5.right = &node9


		got := IsBSTOrdered(&root)
		want := false
		test(t, got, want)
	})
}

func Test_LowestCommonAncestor(t *testing.T) {
	test := func(t testing.TB, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("Find lowest common ancestor 1", func(t *testing.T){
		root := TreeNode{value: 3}
		node1 := TreeNode{value: 5}
		node2 := TreeNode{value: 1}
		node3 := TreeNode{value: 6}
		node4 := TreeNode{value: 2}
		node5 := TreeNode{value: 0}
		node6 := TreeNode{value: 8}
		node7 := TreeNode{value: 7}
		node8 := TreeNode{value: 4}

	 	root.left = &node2; root.right = &node1
		node1.left = &node8; node1.right = &node3
		node2.left = &node5; node2.right = &node4
		node3.right = &node6
		node6.left = &node7

		got := LowestCommonAncestor(&root, 1, 5)
		want := 3
		test(t, got, want)
	})

	t.Run("Find lowest common ancestor 2", func(t *testing.T){
		root := TreeNode{value: 3}
		node1 := TreeNode{value: 5}
		node2 := TreeNode{value: 1}
		node3 := TreeNode{value: 6}
		node4 := TreeNode{value: 2}
		node5 := TreeNode{value: 0}
		node6 := TreeNode{value: 8}
		node7 := TreeNode{value: 7}
		node8 := TreeNode{value: 4}

		root.left = &node2; root.right = &node1
		node1.left = &node8; node1.right = &node3
		node2.left = &node5; node2.right = &node4
		node3.right = &node6
		node6.left = &node7

		got := LowestCommonAncestor(&root, 8, 4)
		want := 5
		test(t, got, want)
	})

	t.Run("Find lowest common ancestor 3", func(t *testing.T){
		root := TreeNode{value: 1}
		node1 := TreeNode{value: 2}

		root.right = &node1

		got := LowestCommonAncestor(&root, 1, 2)
		want := 1
		test(t, got, want)
	})
}

func Test_MergeSortedLists(t *testing.T){
	test := func(t testing.TB, got, want *ListNode) {
		t.Helper()
		if want == nil {
			if got != nil {
				t.Errorf("got %v want %v", got, want)
			}
		} else {
			listGot := new(List)
			listGot.GetIntToList(got)
			listWant := new(List)
			listWant.GetIntToList(want)
			if len(listGot.list) == len(listWant.list) {
				for i, value := range listGot.list{
					if value != listWant.list[i] {
						t.Errorf("got %v want %v", listGot.list, listWant.list)
					}
				}
			} else {
				t.Errorf("got %v want %v", listGot.list, listWant.list)
			}
		}
	}


	t.Run("Merge 2 Sorted Lists", func(t *testing.T){
		//l1 = [1,2,4], l2 = [1,3,4]
		node1 := &ListNode{value:1}
		node2 := &ListNode{value:2}
		node3 := &ListNode{value:4}
		node4 := &ListNode{value:1}
		node5 := &ListNode{value:3}
		node6 := &ListNode{value:4}

		node1.next = node2; node2.next = node3
		node4.next = node5; node5.next = node6

		got := MergeSortedLists(node1, node4)

		//l3 = [1,1,2,3,4,4]
		//node1 -> node4 -> node2 -> node5 -> node3 -> node6
		node1.next = node4; node4.next = node2; node2.next = node5; node5.next = node3; node3.next = node6
		want := node1

		test(t, got, want)
	})

	t.Run("Merge two nil Lists", func(t *testing.T){
		got := MergeSortedLists(nil, nil)
		test(t, got, nil)
	})

	t.Run("Merge 1 list and another nil List", func(t *testing.T){
		node1 := &ListNode{value:1}
		got := MergeSortedLists(node1, nil)
		want := node1
		test(t, got, want)
	})
}

type List struct {
	list []int
}

func (l *List)GetIntToList(node *ListNode){
	if node != nil{
		l.list = append(l.list, node.value)
		l.GetIntToList(node.next)
	}
}

func Test_MthFromEnd(t *testing.T) {
	test := func(t testing.TB, got, want int) {
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("return the last int", func(t *testing.T) {
		//l = [1,2,4]
		node1 := &ListNode{value:1}
		node2 := &ListNode{value:2}
		node3 := &ListNode{value:4}

		node1.next = node2; node2.next = node3

		got := MthFromEnd(node1,0)
		want := 4
		test(t, got, want)
	})

	t.Run("return the second to last int", func(t *testing.T) {
		//l = [1,2,4]
		node1 := &ListNode{value:1}
		node2 := &ListNode{value:2}
		node3 := &ListNode{value:4}

		node1.next = node2; node2.next = node3

		got := MthFromEnd(node1,1)
		want := 2
		test(t, got, want)
	})

}

func Test_StackWithMin(t *testing.T) {
	test := func(t testing.TB, got, want int) {
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("stack with 10, 3, 12, 2", func(t *testing.T) {
		var S = Stack{}
		S.Push(10)
		S.Push(3)
		S.Push(12)
		S.Push(2)

		got := S.Min()
		want := 2
		test(t, got, want)
	})

	t.Run("pop 2 out", func(t *testing.T) {
		var S = Stack{}
		S.Push(10)
		S.Push(3)
		S.Push(12)
		S.Push(2)

		got := S.Pop()
		want := 2
		test(t, got, want)
	})


	t.Run("stack with 10, 3, 12", func(t *testing.T){
		var S = Stack{}
		S.Push(10)
		S.Push(3)
		S.Push(12)
		S.Push(2)
		S.Pop()

		got := S.Min()
		want := 3
		test(t, got, want)
	})

	t.Run("pop 12 out", func(t *testing.T) {
		var S = Stack{}
		S.Push(10)
		S.Push(3)
		S.Push(12)
		S.Push(2)
		S.Pop()

		got := S.Pop()
		want := 12
		test(t, got, want)
	})

	t.Run("stack with 10, 3", func(t *testing.T){
		var S = Stack{}
		S.Push(10)
		S.Push(3)
		S.Push(12)
		S.Push(2)
		S.Pop()
		S.Pop()

		got := S.Min()
		want := 3
		test(t, got, want)
	})
}

func Test_RemoveRowsColumns(t *testing.T) {
	test := func(t testing.TB, got, want [][]int) {
		if len(got) == len(want) {
			for i, array := range got {
				for j, value := range array {
					if value != want[i][j] {
						t.Errorf("got %v want %v", got, want)
					}
				}
			}
		} else {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("test remove rows and columns", func(t *testing.T){
		a := [][]int{{1,2,3},{4,5,6},{7,8,9}}
		got := RemoveRowsColumns(a, 4)
		want := [][]int{{2,3},{8,9}}
		test(t, got, want)
	})

	t.Run("test remove rows and columns", func(t *testing.T){
		a := [][]int{{1,1,1},{1,4,1},{1,1,4}}
		got := RemoveRowsColumns(a, 4)
		want := [][]int{{1}}
		test(t, got, want)
	})

}