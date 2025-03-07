package main

import (
	"fmt"
	"math/big"
	"strings"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Reverse a slice of ints
func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	l1Slice := make([]int, 0)
	l2Slice := make([]int, 0)

	for l1 != nil {
		l1Slice = append(l1Slice, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		l2Slice = append(l2Slice, l2.Val)
		l2 = l2.Next
	}

	l1Slice = reverse(l1Slice)
	l2Slice = reverse(l2Slice)

	//Convert integer array to string
	l1String := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(l1Slice)), ""), "[]")
	l2String := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(l2Slice)), ""), "[]")

	//Convert strings to big integers
	l1Int, l2Int := new(big.Int), new(big.Int)
	l1Int, _ = l1Int.SetString(l1String, 10)
	l2Int, _ = l2Int.SetString(l2String, 10)
	l1Int.Add(l1Int, l2Int)
	sumInt := fmt.Sprintf("%v", l1Int)

	dummy := &ListNode{Val: 0, Next: nil}
	ret := dummy
	// Convert string sum to List return value
	for i := len(sumInt) - 1; i >= 0; i-- {
		dummy.Next = &ListNode{Val: int(sumInt[i]) - 48, Next: nil}
		dummy = dummy.Next
	}

	return ret.Next
}

func main() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

	result := addTwoNumbers(l1, l2)

	for result != nil {
		fmt.Printf("%d -> ", result.Val)
		result = result.Next
	}
}
