package infinitemath

import (
	c "github.com/eightlay/InfiniteMath/iternal/constraints"
	e "github.com/eightlay/InfiniteMath/iternal/errors"
)

// Inner product of two matrices
func Inner[T c.Numeric](m1, m2 *Matrix[T], valsOp, aggOp Operator[T, T]) *Matrix[T] {
	result := ZeroMatrix[T](m1.height, m2.width)
	WriteInner(result, m1, m2, valsOp, aggOp)
	return result
}

// TODO: Apply `by` function to `from` matrix and write result in `to` matrix
func WriteInner[T c.Numeric](to, m1, m2 *Matrix[T], valsOp, aggOp Operator[T, T]) {
	if to.height != m1.height || to.width != m2.width {
		panic(e.ErrWriteInDimensions)
	}

	if m1.width != m2.height {
		panic(e.ErrInnerDimensions)
	}

	for row := uint(0); row < m1.height; row++ {
		for col := uint(0); col < m2.width; col++ {
			toAgg := make([]T, m2.height)

			for k := uint(0); k < m2.height; k++ {
				toAgg[k] = valsOp(m1.Get(row, k), m2.Get(k, col))
			}

			to.vals[row][col] = aggOp(toAgg...)
		}
	}
}
