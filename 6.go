package main

import "fmt"

/*
Sum square difference
Problem 6

The sum of the squares of the first ten natural numbers is,
12 + 22 + ... + 102 = 385

The square of the sum of the first ten natural numbers is,
(1 + 2 + ... + 10)2 = 552 = 3025

Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/

func sumOfSquares(n int) int {
	ans := 0
	for i := 1; i <= n; i++ {
		ans += i * i
	}
	return ans
}

func squareOfSum(n int) int {
	sum := n * (n + 1) / 2
	return sum * sum
}

func main() {
	fmt.Println(squareOfSum(100) - sumOfSquares(100))
}
