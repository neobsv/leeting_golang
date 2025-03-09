package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	current1 := list1
	current2 := list2

	dummy := &ListNode{Val: 0, Next: nil}
	res := dummy

	for current1 != nil && current2 != nil {

		if current1.Val < current2.Val {
			res.Next = current1
			current1 = current1.Next
		} else if current1.Val > current2.Val {
			res.Next = current2
			current2 = current2.Next
		} else {
			res.Next = current1
			current1 = current1.Next
			res = res.Next
			res.Next = current2
			current2 = current2.Next
		}

		res = res.Next

	}

	for current1 != nil {
		res.Next = current1
		current1 = current1.Next
		res = res.Next
	}

	for current2 != nil {
		res.Next = current2
		current2 = current2.Next
		res = res.Next
	}

	return dummy.Next

}

func main() {
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}

	mergedList := mergeTwoLists(list1, list2)

	for mergedList != nil {
		fmt.Printf("%d -> ", mergedList.Val)
		mergedList = mergedList.Next
	}
	fmt.Println("nil")
}
