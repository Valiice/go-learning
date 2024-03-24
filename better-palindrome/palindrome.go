package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome(121))
}

func isPalindrome(x int) bool {
    switch {
    case x < 0:
        return false
    case x >= 1 && x < 10:
        return true
    default:
        revX := 0
        for i := x; i > 0; i /= 10 {
            revX = 10*revX + i%10    
        }
        return revX == x
    }
}
