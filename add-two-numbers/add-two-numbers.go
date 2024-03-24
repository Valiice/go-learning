package main

import (
	"fmt"
)

func main() {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	result := addTwoNumbers(l1, l2)
	fmt.Println(result.Val, result.Next.Val, result.Next.Next.Val)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		l1 = &ListNode{0, nil}
	}
	if l2 == nil {
		l2 = &ListNode{0, nil}
	}

	sum := l1.Val + l2.Val
	carry := sum / 10
	digit := sum % 10
	next := &ListNode{digit, nil}

	if l1.Next == nil && l2.Next == nil {
		if carry > 0 {
			next.Next = &ListNode{carry, nil}
		}
		return next
	}

	l1 = l1.Next
	l2 = l2.Next
	if l1 != nil {
		l1.Val += carry
	} else {
		l2.Val += carry
	}
	next.Next = addTwoNumbers(l1, l2)
	return next
}
