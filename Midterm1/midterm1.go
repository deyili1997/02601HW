// 02-601 Midterm #1, Fall 2021
//
// You have until Friday, October 8th at 11:59pm to complete this exam.  You
// may not discuss the exam with anyone except the course staff.
//
// You may look at any material on golang.org or on the course's Diderot site.
// No other material may be consulted.
//
// There are 7 programming problems. Do not change the function signatures.
// Please leave the problem descriptions in your submitted file.
//
// Please submit a gzipped tar file that contains only one file (midterm1.go)
// that contains the solutions to all the problems.
//
// Each function has a dummy return statement that you should replace with your
// code. Your code should compile, so if you choose to not do a problem, keep
// the an appropriate `return` statement in the function.
//
// You may want to create a `main` function in a separate main.go file that
// tests your functions. Do not turn in such a main.go file, and do not add a
// main() function to this midterm1.go file. You can add helper functions to
// this file if you believe that is the best organization of your code.
//
// You can use any standard packages included with Go.
//
// Each problem is worth 14 points. You get 2 points for correctly submitting
// the file.
//
// NOTE: You will only be able to submit this file ONCE. So be sure you are
// finished before submitting.

// Replace the following with your information:
// Name: Eric Li
// Andrew ID: deyil

package main
import(
	// "fmt"
	"strconv"
)
/*******************************************************************************************
Problem 1. Write a function that takes in a list []bool and returns the number
of times the list switches from true to  false or from false to true when read
left-to-right.

For example, if the list is []bool{true, true, false, false, true, false}, your
function should return 3.
*******************************************************************************************/
func CountSwitches(b []bool) int {
	//initial counter
	counter := 0
	for i := 0; i < len(b)-1;i++ {
//if switch happens, counter + 1
		if b[i] != b[i+1] {
			counter++
		}
	}
	return counter
}
//
/*******************************************************************************************
Problem 2. Write a function that takes in a [][]int list and return the
smallest a integer i such that f[x][i] < 0 for some integer x. If no such i
exists, return -1.

For example, if the input is:
    []int{
        []int{3,4,-5,1,7},
        []int{10,10,10},
        []int{2, 1, 0, -1},
        []int{3,-3,-3,0},
        []int{4,-4,5,0},
    }

Your function should return 1.
*******************************************************************************************/
func LeftmostNegative(f [][]int) int {
	// make a list containing the index of the first item in a row that is < 0
	candidate := make([]int, 0)
	//go through the 2-D matrix
	for row := range f {
		for col := 0; col < len(f[row]); col++ {
			// if find the first item < 0 then jump out of the loop
			if f[row][col]<0 {
				candidate = append(candidate,col)
				break
			}
		}
	}
	// if the item of the matrix < 0 exits
	if len(candidate) > 0 {
		smallest := FindSmallest(candidate)
		return smallest
	}
	//if no such i exists
	return -1
}
// take a list of integer as input and return the smallest item
func FindSmallest(c []int) int {
	smallest := c[0]
	for _,j := range c {
		if j < smallest {
			smallest = j
		}
	}
	return smallest
}
/*******************************************************************************************
Problem 3. Write a function called FirstUnique that takes a slice of integers
and returns the first integer that only occurs once in the slice. If no
integers occur only once, return 0.

Examples:
    FirstUnique([]int{2, 3, 4, 5, 2, 4, 5}) should return 3.
    FirstUnique([]int{7, 8, 2, 8, 2, 7, 8}) should return 0.
    FirstUnique([]int{8, 8, 8, 9, 8, 6, 8, 8}) should return 9.
*******************************************************************************************/
func FirstUnique(list []int) int {
	if len(list) == 0 {
		return 0
	}
	//generate a map of frequency of the item in the input list
	freqTable := make(map[int]int)
	//a list contain the item in the input list that only occur once
	candidate := make([]int,0)
	//make a list of index (the index in the original input list) of the item in candidate list
	candidateIndex := make([]int,0)
	// generate frequency table
	for _,j := range list {
		freqTable[j]++
	}
	// find items that only occur once
	for m,n := range freqTable {
		if n == 1{
			candidate = append(candidate,m)
		}
	}
	//if there is item that only occur once
	if len(candidate) > 0 {
		for p := range list {
			for q := range candidate {
				if candidate[q] == list[p] {
					candidateIndex = append(candidateIndex, p)
				}
			}
		}
		//find the first unique item
		firstCandidate := FindSmallest(candidateIndex)
		return list[firstCandidate]
	}
	//if there is no unique item
	return 0
}

/*******************************************************************************************
Problem 4. A list of numbers A is a circular permutation of another list B if
you can write both lists around a circle and rotate the circles so that the
positions of the numbers are identical.

For example:
    A = 1,7,8,10,31,14
    B = 10,31,14,1,7,8
    C = 7,8,1,10,14,31
A and B are circular permutations of each other, while C is not a circular
permutation of A or B.

Write a function IsCircularPermutation that takes two int lists as parameters
and returns true if one is a circular permutation of the other (and false
otherwise).
*******************************************************************************************/
func IsCircularPermutation(a, b []int) bool {
	if len(a) == 0 || len(b) == 0 {
		return false
	}
	//if the length of a and b are different, then return false
	if len(a) != len(b) {
		return false
	}
	//move every item in a backward by "step" steps
	for step := 0; step < len(a); step++ {
		newOrderList := make([]int,0)
		for i := range a {
			newIndex := i - step
			if newIndex < 0 {
				newIndex = newIndex + len(a)
			}
			newOrderList = append(newOrderList, a[newIndex])
		}
		//if the new ordered list match is identical to b
		if SameSlice(newOrderList,b) {
			return true
		}
	}
	return false
}
//evaluate whether two slices are the same
func SameSlice(a,b []int) bool {
	result := true
	for i := 0; i < len(a); i++ {
		if a[i]!=b[i] {
			result = false
		}
	}
	return result
}
/*******************************************************************************************
Problem 5. Write a function Contains(l1, l2 []int) bool that returns true if
the integers in l2 are a subset of the integers in l1 and false otherwise.
That is, the function returns true if every integer that appears in l2 also
appears someplace in l1. If an integer occurs in l2 more than once, it needs to
occur at least that many times in l1.
*******************************************************************************************/
func Contains(l1, l2 []int) bool {
	//generate a freqTable of l1 and l2
	freq1 := FreqTable(l1)
	freq2 := FreqTable(l2)
	// for each key, value in freq2, check whether they are contained in freq1 and the frequency satisfy the requirement
	resultList := make([]bool,0)
	for k, v := range freq2 {
		result := false
		for m, n := range freq1 {
			if m == k && n >= v {
				result = true
				resultList = append(resultList,result)
				break
			}
		}
		resultList = append(resultList,result)
	}
	//check whether there is any key in freq2 does not satisfy the requirement
	for _,j := range resultList {
		if j != true {
			return false
		}
	}
	return true
}


func FreqTable(l []int) map[int]int {
	table := make(map[int]int)
	for _,j := range l {
		table[j]++
	}
	return table
}

/*******************************************************************************************
Problem 6. Write a function SquareNumbers that takes a list of integers and
returns the numbers in that list that are square numbers (can be written as x*x
for some integer x). The output list should be in the same order as the input list.

Example: If your input is []int{1,5,9,9,20} the returned value should be []int{1,9,9}.
*******************************************************************************************/
func SquareNumbers(a []int) []int {
	if len(a)== 0 {
		return []int{}
	}
	//generate the result list
	result := make([]int,0)
	//step1: wipe out the items that cannot be a suqare number according to maths
	cleanedList := WipeOut(a)
	//step2: sort the remaining items
	orderedList := SortList(cleanedList)
	//step3: use MinBase and MaxBase to narrow down the loop rounds
	MaxBase := FindBase(orderedList[len(orderedList)-1])+1
	MinBase := FindBase(orderedList[0])
	//step4: find the square numbers
	for i := MinBase; i <= MaxBase; i++ {
		for _,j := range orderedList {
			if j == i * i {
				result = append(result,j)
			}
		}
	}
	return result
}
//wipe out item that its last digit is not 0,1,4,5,6,9 according to the property of SquareNumbers
func WipeOut(a []int) []int {
	newList := make([]int,0)
	for _,j := range a {
		strInt := strconv.Itoa(j)
		digit := string(strInt[len(strInt)-1])
		if digit == "0" || digit == "1" || digit == "4" || digit == "5" || digit == "6" ||digit == "9" {
			newList = append(newList,j)
		}
	}
	return newList
}
// bubble sort the list
func SortList(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := i+1; j< len(a) - i; j++ {
			if a[j] < a[j-1] {
				a[j],a[j-1] = a[j-1],a[j]
			}
		}
	}
	return a
}
// find a base that base * base < input and (base + 1)* (base + 1) > input
func FindBase(b int) int {
	var base int
	for i:=0; i<b; i++{
		if i * i <= b && (i+1) * (i+1) > b {
			base = i
			break
		}
	}
	return base
}

/*******************************************************************************************
Problem 7. Consider the following Node type. It contains a single field which
is a pointer to another Node. That other Node of course will itself contain a
field pointing to a Node.  By following these pointers we can walk through a
series of Nodes. If `next` is nil, the series of Nodes ends.

Write a function HasCycle that walks through the series of Nodes staarting with
the parameter n and returns true if you can ever return back to a node you have
already visited, and false otherwise.

Be sure to consider the following case:

n -> n1 -> n2 -> n3 -> n4 -> n5 -+
                 ^               |
                 |               |
                 +---------------+

which should return true.

You cannot modify the Node type or the pointers in the Node.
*******************************************************************************************/

type Node struct {
	next *Node
}

func HasCycle(n *Node) bool {
	if n == nil || n.next == nil {
		return false
	}
	link := make([]*Node,0)
	for n.next != nil {
		link = append(link,n)
		n = n.next
		if len(link) >= 2 {
			for i:=0; i<len(link)-1;i++{
				if link[len(link)-1] == link[i] {
					return true
				}
			}
		}
	}
	return false
}

// END OF MIDTERM
