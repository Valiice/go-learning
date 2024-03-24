package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(0))
}

func isPalindrome(i int) bool {
	return i == reverseInt(i)
}

func reverseInt(x int) int {
	if x < 0 {
		return -x
	}
	sum := 0
	for x != 0 {
		sum = sum*10 + x%10
		x /= 10
	}
	return sum
}
