package main

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	// Test addTwoNumbers function
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	result := addTwoNumbers(l1, l2)
	if result.Val != 7 || result.Next.Val != 0 || result.Next.Next.Val != 8 {
		t.Errorf("Expected 7 -> 0 -> 8, but got %d -> %d -> %d", result.Val, result.Next.Val, result.Next.Next.Val)
	}
}

func TestAddTwoNumbersAnotherOne(t *testing.T) {
	// Test addTwoNumbers function
	l1 := &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}
	l2 := &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}
	result := addTwoNumbers(l1, l2)
	if result.Val != 8 || result.Next.Val != 9 || result.Next.Next.Val != 9 || result.Next.Next.Next.Val != 9 || result.Next.Next.Next.Next.Val != 0 || result.Next.Next.Next.Next.Next.Val != 0 || result.Next.Next.Next.Next.Next.Next.Val != 0 || result.Next.Next.Next.Next.Next.Next.Next.Val != 1 {
		t.Errorf("Expected 8 -> 9 -> 9 -> 9 -> 0 -> 0 -> 0 -> 1, but got %d -> %d -> %d -> %d -> %d -> %d -> %d -> %d", result.Val, result.Next.Val, result.Next.Next.Val, result.Next.Next.Next.Val, result.Next.Next.Next.Next.Val, result.Next.Next.Next.Next.Next.Val, result.Next.Next.Next.Next.Next.Next.Val, result.Next.Next.Next.Next.Next.Next.Next.Val)
	}
}
