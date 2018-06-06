/*
Package diffsquares finds the difference between the square
of the sum and the sum of the squares of the first N natural
numbers.

The square of the sum of the first ten natural numbers is
(1 + 2 + ... + 10)² = 55² = 3025.

The sum of the squares of the first ten natural numbers is
1² + 2² + ... + 10² = 385.

Hence the difference between the square of the sum of the first
ten natural numbers and the sum of the squares of the first ten
natural numbers is 3025 - 385 = 2640.
*/
package diffsquares

// SquareOfSums returns the square of sums from 1 to the number n.
func SquareOfSums(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

// SumOfSquares returns the sum of squares from 1 to the number n.
func SumOfSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

// Difference returns the difference between the square
// of the sum and the sum of the squares of the first N natural
// numbers.
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
