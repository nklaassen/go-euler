package main

import "fmt"

/*
Largest palindrome product
Problem 4

A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func hasThreeDigitFactors(n int) bool {
	for i := max(n/999, 100); i <= min(n/100, 999); i++ {
		if n%i == 0 {
			return true
		}
	}
	return false
}

func main() {
	// enumerate the 6-digit palindromes in decreasing order, and check if they have 3 digit factors
	for outer := 9; outer >= 0; outer-- {
		for middle := 9; middle >= 0; middle-- {
			for inner := 9; inner >= 0; inner-- {
				n := outer*100001 + middle*10010 + inner*1100
				if hasThreeDigitFactors(n) {
					fmt.Println(n)
					return
				}
			}
		}
	}
}
