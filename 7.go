package main

import "fmt"

/*
10001st prime
Problem 7

By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?
*/

func log2(n int) int {
	for i := uint(31); i >= 0; i-- {
		if n&(1<<i) != 0 {
			return int(i)
		}
	}
	return 0
}

func main() {
	n := 10001

	// use the sieve of eratosthenes
	size := 2 * n * log2(n)          // this is a guess, turns out to be okay
	composites := make([]bool, size) // keep a slice of composite numbers to take advantage of zero-values
	numPrimes := 0
	for i := 2; i < size; i++ {
		if !composites[i] {
			numPrimes++
			if numPrimes == n {
				fmt.Println(i)
				return
			}
			for j := i * i; j < size; j += i {
				composites[j] = true
			}
		}
	}
}
