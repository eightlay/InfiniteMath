package infinitemath

import (
	c "github.com/eightlay/InfiniteMath/iternal/constraints"
)

func Dot[T c.Numeric](m1, m2 *Matrix[T]) *Matrix[T] {
	return Inner(m1, m2, multOp[T], addOp[T])
}
