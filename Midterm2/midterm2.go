// 02-601 Midterm #2, Fall 2021
//
// You have until Monday, November 8th at 11:59pm to complete this exam.  You
// may not discuss the exam with anyone except the course staff.
//
// You may look at any material on golang.org or on the course's Diderot site.
// No other material may be consulted.
//
// There are 6 programming problems. Do not change the function signatures.
// Please leave the problem descriptions in your submitted file.
//
// Please submit a gzipped tar file that contains only one file (midterm2.go)
// that contains the solutions to all the problems.
//
// Each function has a dummy return statement that you should replace with your
// code. Your code should compile, so if you choose to not do a problem, keep
// the appropriate `return` statement in the function.
//
// You may want to create a `main` function in a separate main.go file that
// tests your functions. Do not turn in such a main.go file, and do not add a
// main() function to this midterm2.go file. You can add helper functions to
// this file if you believe that is the best organization of your code.
//
// You can use any standard packages included with Go.
//
// Each problem is worth 16 points. You get 4 points for correctly submitting
// the file.
//
// NOTE: You will only be able to submit this file ONCE. So be sure you are
// finished before submitting.

// Replace the following with your information:
// Name: Eric Li
// Andrew ID: deyil

package main

/*******************************************************************************************
Problem 1. Write the function "IsBSTOrdered()": If the values in the binary
tree rooted at `root` satisfy the binary search tree ordering, return true. If
the values do NOT all satisfy the binary search tree ordering, return false.
You may assume there are no duplicate items in the tree.
*******************************************************************************************/

// TreeNode is used in this and the next problem. It represents a node in a binary tree.
type TreeNode struct {
	left  *TreeNode // subtree with smaller values (nil if no subtree)
	right *TreeNode // subtree with larger values (nil if no subtree)
	value int
}

func IsBSTOrdered(root *TreeNode) bool {
	if root.left != nil && root.right != nil {
		left := FindChildren(root.left)
		right := FindChildren(root.right)
		//check BST ordering
		for _, l := range left {
			if l.value >= root.value {
				return false
			}
		}
		for _, r := range right {
			if r.value <= root.value {
				return false
			}
		}
		return IsBSTOrdered(root.left) && IsBSTOrdered(root.right)
		//if there is no right branch
	} else if root.left != nil && root.right == nil {
		left := FindChildren(root.left)
		for _, l := range left {
			if l.value >= root.value {
				return false
			}
		}
		return IsBSTOrdered(root.left)
		//if there is no left branch
	} else if root.left == nil && root.right != nil {
		right := FindChildren(root.right)
		for _, r := range right {
			if r.value <= root.value {
				return false
			}
		}
		return IsBSTOrdered(root.right)
		//if this is a leaf
	} else if root.left == nil && root.right == nil {
		return true
	}
	return false
}

//Find all the children node under the node
func FindChildren(root *TreeNode) []*TreeNode {
	var children []*TreeNode
	children = append(children, root)
	if root.left != nil && root.right != nil {
		children = append(children, FindChildren(root.left)...)
		children = append(children, FindChildren(root.right)...)
	} else if root.left != nil && root.right == nil {
		children = append(children, FindChildren(root.left)...)
	} else if root.left == nil && root.right != nil {
		children = append(children, FindChildren(root.right)...)
	} else if root.left == nil && root.right == nil {
		return children
	}
	return children
}

/*******************************************************************************************
Problem 2. Lowest common ancestor. Given x and y, return the value in the tree
rooted at `root` that is in the node that is the lowest common ancestor of x
and y. The ancestors of a node x are all the nodes between x and the root. The
lowest common ancestor of two nodes x and y is the node that is an ancestor of
both x and y that is closest to x and y. You may assume nodes for x and y exist
in the tree, and that there are no nodes that share the same value in the tree.

Hint: Think about the BST tree ordering.
*******************************************************************************************/

func LowestCommonAncestor(root *TreeNode, x, y int) int {
	path1 := FindPath(root, x)
	path2 := FindPath(root, y)
	for _, n := range path1 {
		for _, k := range path2 {
			if n == k {
				return n
			}
		}
	}
	panic("You should not go here")
	return -1
}

func FindPath(root *TreeNode, t int) []int {
	path := make([]int, 0)
	if root.value == t {
		// fmt.Println("root.value == t")
		path = append(path, root.value)
	} else {
		// fmt.Println("root.value != t")
		if root.left != nil && root.right != nil {
			// fmt.Println("root.left != nil && root.right != nil")
			left := FindPath(root.left, t)
			right := FindPath(root.right, t)
			if len(left) == 0 && len(right) != 0 {
				path = append(path, right...)
				path = append(path, root.value)
			} else if len(left) != 0 && len(right) == 0 {
				// fmt.Println("find left")
				path = append(path, left...)
				path = append(path, root.value)

			} else if len(left) == 0 && len(right) == 0 {
				// fmt.Println("not find left nor right")
				return path
			}
		} else if root.left != nil && root.right == nil {
			// fmt.Println("root.left != nil && root.right == nil")
			left := FindPath(root.left, t)
			if len(left) != 0 {
				path = append(path, left...)
				path = append(path, root.value)
			}
		} else if root.left == nil && root.right != nil {
			// fmt.Println("root.left == nil && root.right != nil")
			right := FindPath(root.right, t)
			if len(right) != 0 {
				path = append(path, right...)
				path = append(path, root.value)
			}
		} else if root.left == nil && root.right == nil {
			// fmt.Println("root.left == nil && root.right == nil")
			return path
		}
	}
	return path
}

/*******************************************************************************************
Problem 3. Sorted linked list merge. You are given two linked lists, l1 and l2,
each of which contain values that are sorted by non-decreasing values. Return a
new linked list that contains all of the nodes of l1 and l2 and that is sorted.
You should not allocate new nodes; reuse the nodes in the input lists.

That is if l1 contains n items and l2 contains m items, the list you return
will have n+m items in sorted order.
*******************************************************************************************/

// ListNode is used in this problem and the next one. It is a node in a linked list.
type ListNode struct {
	value int
	next  *ListNode // next node, nil if at end.
}

func MergeSortedLists(l1, l2 *ListNode) *ListNode {
	//to judge the nil situation
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1 == nil && l2 == nil {
		return nil
	}
	// to judge the l1 l2 length are 1 without interfering with the recursion out case
	if l1.next == nil && l2.next == nil {
		if l1.value < l2.value {
			l1.next = l2
			return l1
		} else if l1.value > l2.value {
			l2.next = l1
			return l2
		} else if l1.value == l2.value {
			l1.next = l2
			return l1
		}
	}
	// recursion out case
	if l1.next == nil && l2.next != nil && l1.value <= l2.value {
		l1.next = l2
		return l1
	}
	//main body
	if l1.value < l2.value {
		next := FindNext(l1, l2)
		keepinfo := next.next
		next.next = l2
		MergeSortedLists(next.next, keepinfo)
	} else if l1.value > l2.value {
		next := FindNext(l2, l1)
		keepinfo := next.next
		next.next = l1
		MergeSortedLists(next.next, keepinfo)
	} else if l1.value == l2.value {
		keepinfo := l1.next
		l1.next = l2
		MergeSortedLists(l1.next, keepinfo)
	}

	if l1.value > l2.value {
		return l2
	} else if l1.value < l2.value {
		return l1
	} else if l1.value == l2.value {
		return l1
	}
	return nil
}

//find next node to connect
func FindNext(l1, l2 *ListNode) *ListNode {
	var next *ListNode
	if l1.next == nil {
		return l1
	} else {
		if l1.next.value < l2.value {
			next = FindNext(l1.next, l2)
		} else {
			return l1
		}
	}
	return next
}

/*******************************************************************************************
Problem 4. Given an UNSORTED linked list, and an integer m, return the integer
(value) that is at position m from the END of the linked list. That is if m ==
0, return the last item in the linked list; if m == 1, return the second to
last, etc. You can assume that m is less than the length of the linked list.
*******************************************************************************************/

func MthFromEnd(ll *ListNode, m int) int {
	length := FindLength(ll)
	value := GoToTheNode(ll, length-m)
	return value
}

func FindLength(n *ListNode) int {
	length := 0
	if n.next != nil {
		length = 1 + FindLength(n.next)
	} else {
		length = 1
	}
	return length
}

func GoToTheNode(n *ListNode, k int) int {
	if k == 1 {
		return n.value
	}
	value := GoToTheNode(n.next, k-1)
	return value
}

/*******************************************************************************************
Problem 5. Stack with Min(). Modify the following Stack data structure to
support the following operations:

	Push(): push an item on the stack
	Pop(): pop an item from the stack
	Min(): return the value of the smallest integer on the stack (WITHOUT modifying the
		   items on the stack)

Each of the above operations must run in time *independent* of the number of
items on the stack. That is they should take O(1) time. Consequently, you
should not create any maps or arrays. You will have to modify the StackItem
and/or the Stack types and the Push and Pop operations and implement the Min
operation.

Note: you can create a new Stack with "var S = Stack{}" (since everything
defaults to nil); keep this approach to creating a new stack.

Example: The following code should work:

		var S = Stack{}
		S.Push(10)
		S.Push(3)
		S.Push(12)
		S.Push(2)
		fmt.Println(S.Min())	// prints 2
		S.Pop()
		fmt.Println(S.Min())	// prints 3
		S.Pop()
		fmt.Println(S.Min())	// prints 3

Hint: as one Push()es new items, the min can only go down.
*******************************************************************************************/

type Stack struct {
	top *StackItem // pointer to the item on the top of the stack
}

type StackItem struct {
	prev  *StackItem // pointer to the next item on the stack
	value int
	min   int
}

func CreatItem(pre *StackItem, v, m int) *StackItem {
	return &StackItem{
		prev:  pre,
		value: v,
		//store the min value among all the layers under this layer
		min: m,
	}
}

// Push adds a new item to the top of the stack. It runs in time independent
// of the number of elements in the stack.
func (s *Stack) Push(v int) {
	if s.top == nil {
		s.top = CreatItem(nil, v, v)
	} else {
		if v < s.top.min {
			min := v
			s.top = CreatItem(s.top, v, min)
		} else {
			min := s.top.min
			s.top = CreatItem(s.top, v, min)
		}
	}
}

// Pop removes and returns the item at the top of the Stack. It runs in
// time independent of the number of elements in the Stack.
func (s *Stack) Pop() int {
	if s.top == nil {
		panic("Pop on an empty stack!")
	}
	v := s.top.value
	s.top = s.top.prev
	return v
}

// Min returns the smallest integer on the stack without changing the items on the stack.
// It runs in time independent of the number of items in the Stack.
func (s *Stack) Min() int {
	// MODIFY THE ABOVE Push and Pop FUNCTIONS and the DATA STRUCTURES
	// WRITE THIS FUNCTION
	// YOUR Min FUNCTION SHOULD RUN IN TIME *INDEPENDENT* OF THE NUMBER OF ITEMS IN THE STACK
	return s.top.min
}

/*******************************************************************************************
Problem 6. Knock out number. Given a (possibly non-square) matrix `matrix` and
a number x.  Remove all rows and columns that contain x and return the smaller
matrix. For example, if x appears in cell (4,5) and (4,8) then remove row 4 and
columns 5 and 8. Rows and columns start at index 0. Your function can "destroy"
the input matrix.
*******************************************************************************************/

func RemoveRowsColumns(matrix [][]int, x int) [][]int {
	rowsDelete := make([]int, 0)
	colsDelete := make([]int, 0)
	//find the index of rows and cols to be deleted
	for r := range matrix {
		for c, e := range matrix[r] {
			if e == x {
				if !AlreadyIn(rowsDelete, r) {
					rowsDelete = append(rowsDelete, r)
				}
				if !AlreadyIn(colsDelete, c) {
					colsDelete = append(colsDelete, c)
				}
			}
		}
	}
	//sort the index list waiting to be deleted
	SortedRow := SelectionSort(rowsDelete)
	SortedCol := SelectionSort(colsDelete)
	for r := len(SortedRow) - 1; r >= 0; r-- {
		matrix = append(matrix[:SortedRow[r]], matrix[(SortedRow[r]+1):]...)
	}
	for r := range matrix {
		for c := len(SortedCol) - 1; c >= 0; c-- {
			matrix[r] = append(matrix[r][:SortedCol[c]], matrix[r][(SortedCol[c]+1):]...)
		}
	}

	return matrix
}

//if the index is already recorded
func AlreadyIn(m []int, n int) bool {
	for _, k := range m {
		if k == n {
			return true
		}
	}
	return false
}

func SelectionSort(inList []int) []int {
	outList := make([]int, 0)
	n := len(inList)
	for i := 1; i <= n; i++ {
		// select the minimum element left in our inList
		j, m := Min(inList)
		outList = append(outList, m)
		//remove minimum from inList
		inList = append(inList[:j], inList[j+1:]...)
	}

	return outList
}

func Min(inList []int) (int, int) {
	// return minimum index and element of a list
	index := 0
	min := inList[0]
	for i, m := range inList {
		if m < min {
			index = i
			min = m
		}
	}

	return index, min

}

// END OF MIDTERM
