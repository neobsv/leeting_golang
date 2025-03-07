package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Remove the nth node from the end of the list, counting backwards from the tail node.
func removeNthFromEndOld(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	// nth node from the tail node is the len(list) - n - 1 node from the head

	// Find length of the list
	var length int = 1
	var current *ListNode = head
	for current.Next != nil {
		current = current.Next
		length++
	}

	var nthNode = length - n - 1

	// Edge case where there is only one node in the list
	if nthNode < 0 {
		return nil
	}

	// Move to the nth node and skip past it connecting nodes n-1 and n+1
	current = head
	for nthNode > 0 && current.Next != nil {
		current = current.Next
		nthNode--
	}

	if current.Next != nil {
		current.Next = current.Next.Next
	}

	return head
}

// nth node from the tail node is the len(list) - n - 1 node from the head
func removeNthNodeFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	// Find length of the list
	var length int = 0
	var current *ListNode = head
	for current != nil {
		current = current.Next
		length++
	}

	// Edge case where there is only one node in the list
	if length == n {
		return head.Next
	}

	// Move to the nth node and skip past it connecting nodes n-1 and n+1
	current = head
	for current != nil {
		if length == n+1 {
			current.Next = current.Next.Next
			break
		}
		current = current.Next
		length--
	}

	return head
}

func main() {
	// Create linked list: 1 -> 2 -> 3 -> 4 -> 5
	head := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}

	head.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	// Remove the 2nd node from the end
	head = removeNthNodeFromEnd(head, 2)

	// Print the modified linked list
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")

	head1 := &ListNode{Val: 1}
	head1.Next = nil

	head1 = removeNthNodeFromEnd(head1, 1)

	// Print the modified linked list
	current = head1
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")

}
