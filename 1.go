package main

import "fmt"

/*
Multiples of 3 and 5
Problem 1

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/

func sumOfMultiplesLessThan(k, n int) int {
	// find the number of multiples of k less than n
	numMultiples := (n - 1) / k
	// return the sum of all multiples of k less than n
	// this is calculated by multiplying the sum of all integers up to and including numMultiples by k
	return k * numMultiples * (numMultiples + 1) / 2
}

func main() {
	n := 1000
	// a problem of this size could trivially be solved with a for loop, but that's no fun
	// add the sum of all multiples of 3 and 5 less than n, and then subtract the sum of their common multiples
	ans := sumOfMultiplesLessThan(3, n) + sumOfMultiplesLessThan(5, n) - sumOfMultiplesLessThan(15, n)
	fmt.Println(ans)
}
