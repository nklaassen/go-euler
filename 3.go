package main

import (
	"fmt"
)

/*
Largest prime factor
Problem 3

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/

func smallestFactor(n int) int {
	// not terribly efficient, but easy to write
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return i
		}
	}
	return n
}

func largestPrimeFactor(n int) int {
	// again, not terribly efficient, but easy to write
	x := smallestFactor(n)
	if x == n {
		return n
	}
	return largestPrimeFactor(n / x)
}

func main() {
	fmt.Println(largestPrimeFactor(600851475143))
}
