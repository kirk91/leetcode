// source: https://leetcode.com/problems/merge-two-sorted-lists/

/*
Merge two sorted linked lists and return it as a new list.
The new list should be made by splicing together the nodes of the first two lists.
*/

package main

type ListNode struct{
	Val int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }

    var head, p *ListNode
    if l1.Val < l2.Val {
        head = l1
        p = l1
        l1 = l1.Next
    }else {
        head = l2
        p = l2
        l2 = l2.Next
    }
    for ; l1 != nil ; l1 = l1.Next {
        for ; l2 != nil; l2 = l2.Next {
            if l2.Val > l1.Val {
                break
            }
            p.Next = l2
            p = l2
        }
        p.Next = l1
        p = l1
    }

    if l2 != nil {
        p.Next = l2
    }

    return head
}
