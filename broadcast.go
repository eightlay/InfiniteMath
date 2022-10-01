package infinitemath

import (
	c "github.com/eightlay/InfiniteMath/iternal/constraints"
	e "github.com/eightlay/InfiniteMath/iternal/errors"
)

func Broadcast[T c.Numeric](to *Matrix[T], from *Matrix[T], by Operator[T]) {
	if to.height != from.height || to.width != from.width {
		panic(e.ErrDimensions)
	}

	for i := uint(0); i < to.height; i++ {
		for j := uint(0); j < to.width; j++ {
			to.vals[i][j] = by(to.vals[i][j], from.vals[i][j])
		}
	}
}
