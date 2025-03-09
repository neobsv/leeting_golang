package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {

	dummy := &ListNode{Val: 0, Next: head}

	prev := dummy
	current := head

	for current != nil {

		for current.Next != nil && current.Val == current.Next.Val {
			current = current.Next
		}

		// no duplicates case, the id of the next node for prev is the
		// current node, that means the above for loop didn't run, so no duplicates
		if prev.Next == current {
			prev = current
		} else {
			// skip over duplicates, the above for loop ran and we want to skip all the duplicates
			// between prev and current, current pointer is at the last duplicate now
			prev.Next = current.Next
		}

		current = current.Next
	}

	return dummy.Next

}

func main() {
	// Create a linked list: 1->2->3->3->4->4->5
	node5 := &ListNode{Val: 5, Next: nil}
	node4 := &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: node5}}
	node3 := &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: node4}}
	node2 := &ListNode{Val: 2, Next: node3}
	node1 := &ListNode{Val: 1, Next: node2}

	// Call the function and print the result
	result := deleteDuplicates(node1)
	for result != nil {
		fmt.Printf("%d -> ", result.Val)
		result = result.Next
	}
	fmt.Println("nil")
}
