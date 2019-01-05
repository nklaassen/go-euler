package main

import "fmt"

/*
Smallest multiple
Problem 5

2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

func main() {
	// the solution should be the product of the prime factors of the numbers 1 to 20, each taken to the greatest power
	//   found in the factorization. Since this example is small it's easier to do much of it manually.
	//primes := []int{2, 3, 5, 7, 11, 13, 17, 19}
	//powers := []int{4, 2, 1, 1, 1, 1, 1, 1}
	ans := 16 * 9 * 5 * 7 * 11 * 13 * 17 * 19
	fmt.Println(ans)
}
