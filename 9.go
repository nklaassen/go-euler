package main

import "fmt"

/*
Special Pythagorean triplet
Problem 9

A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
a^2 + b^2 = c^2

For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/

func main() {
	const n = 1000
	// The Simplest solution is to search all natural numbers a < b < c which sum to n for a Pythagorean triplet.
	// This solution is O(n^3) but it should be fine for such a small input size.
	for a := 1; a <= n-(a+1)-(a+2); a++ {
		for b := a + 1; b < (n-a)/2; b++ {
			c := n - a - b
			if a*a+b*b == c*c {
				fmt.Println(a * b * c)
				return
			}
		}
	}
	fmt.Println("solution not found")
}
