package algorithms

import (
	"leetcode/packages/linkedlist"
	"testing"
)

func TestUnfoldLinkedList(t *testing.T) {
	var head linkedlist.ListNode

	head = linkedlist.ListNode{
		Val: 1,
	}
	if node := UnfoldLinkedList(&head); node.Val != 1 {
		t.Errorf("UnfoldLinkedList expected 1, got: %t", node.Val)
	}

	head = linkedlist.ListNode{
		Val: 1,
		Next: &linkedlist.ListNode{
			Val: 2,
			Next: &linkedlist.ListNode{
				Val: 3,
			},
		},
	}
	if node := UnfoldLinkedList(&head); node.Val != 3 {
		t.Errorf("UnfoldLinkedList expected 3, got: %t", node.Val)
	}

	head = linkedlist.ListNode{
		Val: 1,
		Next: &linkedlist.ListNode{
			Val: 2,
		},
	}
	node := UnfoldLinkedList(&head)
	if node.Val != 2 {
		t.Errorf("UnfoldLinkedList expected 2, got: %t", node.Val)
	}

	if node.Next.Val != 1 {
		t.Errorf("UnfoldLinkedList expected 1, got: %t", node.Next.Val)
	}

}
