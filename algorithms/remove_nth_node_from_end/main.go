// source:https://leetcode.com/problems/remove-nth-node-from-end-of-list/

/*
Given a linked list, remove the nth node from the end of list and return its head.

For example,

   Given linked list: 1->2->3->4->5, and n = 2.

      After removing the second node from the end, the linked list becomes 1->2->3->5.
	  Note:
	  Given n will always be valid.
	  Try to do this in one pass.
*/

package main

import "github.com/kirk91/leetcode/assert"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	var i int
	for i = 0; i < n+1 && cur != nil; i++ {
		cur = cur.Next
	}
	if i == n {
		return head.Next
	}

	pre := head
	for cur != nil {
		pre = pre.Next
		cur = cur.Next
	}
	pre.Next = pre.Next.Next
	return head
}

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	assert.MustEqual(&ListNode{Val: 1}, removeNthFromEnd(head, 1))
}
