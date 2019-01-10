package main

import "fmt"

/*
Summation of primes
Problem 10

The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/

func main() {
	const n = 2000000

	// use the sieve of eratosthenes to find primes
	sum := 0
	var composites [n]bool // keep a slice of composite numbers to take advantage of zero-values
	for i := 2; i < n; i++ {
		if !composites[i] {
			sum += i
			for j := i * i; j < n; j += i {
				composites[j] = true
			}
		}
	}
	fmt.Println(sum)
}
