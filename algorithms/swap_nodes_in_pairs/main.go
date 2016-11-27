// source: https://leetcode.com/problems/swap-nodes-in-pairs/

/*
Given a linked list, swap every two adjacent nodes and return its head.

For example,
Given 1->2->3->4, you should return the list as 2->1->4->3.

Your algorithm should use only constant space. You may not modify the values in the list, only nodes itself can be changed.
*/

package main

import "github.com/kirk91/leetcode/assert"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := new(ListNode)
	dummy.Next = head
	prev, cur := dummy, head
	for cur != nil && cur.Next != nil {
		prev.Next = cur.Next
		cur.Next = cur.Next.Next
		prev.Next.Next = cur
		prev = cur
		cur = cur.Next
	}
	return dummy.Next
}

func main() {
	assert.MustEqual(&ListNode{2, &ListNode{1, &ListNode{3, nil}}},
		swapPairs(&ListNode{1, &ListNode{2, &ListNode{3, nil}}}))
	assert.MustEqual(&ListNode{2, &ListNode{1, &ListNode{4, &ListNode{3, nil}}}},
		swapPairs(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}))
}
