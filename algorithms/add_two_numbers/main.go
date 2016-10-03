// source: https://leetcode.com/problems/add-two-numbers/

/*
You are given two linked lists representing two non-negative numbers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
*/

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	return fmt.Sprintf("%d %s", l.Val, l.Next)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		l, cur *ListNode
		n1     = l1
		n2     = l2

		v1, v2, sum, carry int
	)

	for {
		if n1 == nil && n2 == nil && carry == 0 {
			break
		}

		if n1 == nil {
			v1 = 0
		} else {
			v1 = n1.Val
		}
		if n2 == nil {
			v2 = 0
		} else {
			v2 = n2.Val
		}

		sum = v1 + v2 + carry
		carry = sum / 10
		n := &ListNode{Val: sum % 10}
		if l == nil {
			l = n
		} else {
			cur.Next = n
		}
		cur = n

		if n1 != nil {
			n1 = n1.Next
		}
		if n2 != nil {
			n2 = n2.Next
		}
	}
	return l
}

func main() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	fmt.Println(addTwoNumbers(l1, l2))
}
