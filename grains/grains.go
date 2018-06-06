/*
Package grains calculate the number of grains of wheat on a chessboard given
that the number on each square doubles.

There once was a wise servant who saved the life of a prince. The king
promised to pay whatever the servant could dream up. Knowing that the
king loved chess, the servant told the king he would like to have grains
of wheat. One grain on the first square of a chess board. Two grains on
the next. Four on the third, and so on.

There are 64 squares on a chessboard.
*/
package grains

import (
	"fmt"
)

// Define minimum and maximum values for the input.
const (
	MinValue = 1
	MaxValue = 64
)

// Square returns the number of grains in a specific chessboard square.
func Square(n int) (uint64, error) {
	if n < MinValue || n > MaxValue {
		err := fmt.Errorf("number must be between %v and %v", MinValue, MaxValue)
		return 0, err
	}

	if n == 1 {
		return 1, nil
	}

	total, _ := Square(n - 1)
	return total * 2, nil
}

// Total returns the total number of grains in the whole chessboard.
func Total() uint64 {
	var total uint64
	for i := 1; i <= MaxValue; i++ {
		s, _ := Square(i)
		total += s
	}
	return total
}
