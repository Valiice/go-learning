package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	// Test isPalindrome function
	if !isPalindrome(121) {
		t.Errorf("Expected true, but got false")
	}
	if isPalindrome(-121) {
		t.Errorf("Expected false, but got true")
	}
	if isPalindrome(10) {
		t.Errorf("Expected false, but got true")
	}
}
