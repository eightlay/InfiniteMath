package infinitemath

import (
	c "github.com/eightlay/InfiniteMath/iternal/constraints"
	e "github.com/eightlay/InfiniteMath/iternal/errors"
)

// Broadcast the operator to pairs of corresponding elements of two matrices
func Broadcast[T c.Numeric](to, from *Matrix[T], by Operator[T]) {
	if to.height != from.height || to.width != from.width {
		panic(e.ErrBroadcastDimensions)
	}

	for row := uint(0); row < to.height; row++ {
		for col := uint(0); col < to.width; col++ {
			to.vals[row][col] = by(to.vals[row][col], from.vals[row][col])
		}
	}
}
