package main

import (
	"bytes"
	"fmt"

	"github.com/kirk91/leetcode/assert"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("%d", l.Val))
	if l.Next != nil {
		b.WriteString(fmt.Sprintf(" -> %s", l.Next))
	}
	return b.String()
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	prev, cur := dummy, head
	var i int
	for cur != nil {
		i++
		next := cur.Next
		if i%k == 0 {
			prev = reverse(prev, next)
		}
		cur = next
	}
	return dummy.Next
}

func reverse(prev, end *ListNode) *ListNode {
	head := prev.Next
	cur := head.Next
	for cur != end {
		next := cur.Next
		cur.Next = prev.Next
		prev.Next = cur
		cur = next
	}
	head.Next = end
	return head
}

func main() {
	assert.MustEqual(&ListNode{2, &ListNode{1, nil}}, reverseKGroup(&ListNode{1, &ListNode{2, nil}}, 2))
}
