package algorithms

import (
	"leetcode/packages/linkedlist"
)

func UnfoldLinkedList(head *linkedlist.ListNode) *linkedlist.ListNode {
	var node = head
	var prev *linkedlist.ListNode

	for node.Next != nil {
		temp := node.Next
		node.Next = prev
		prev = node
		node = temp
	}

	if prev != nil {
		node.Next = prev
	}

	return node
}
