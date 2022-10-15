package infinitemath

import (
	c "github.com/eightlay/InfiniteMath/iternal/constraints"
	e "github.com/eightlay/InfiniteMath/iternal/errors"
)

func Inner[T c.Numeric](m1, m2 *Matrix[T], valsOp, aggOp Operator[T]) *Matrix[T] {
	if m1.width != m2.height {
		panic(e.ErrInnerDimensions)
	}

	result := make([][]T, m1.height)

	for row := uint(0); row < m1.height; row++ {
		result[row] = make([]T, m2.width)

		for col := uint(0); col < m2.width; col++ {
			toAgg := make([]T, m2.height)

			for k := uint(0); k < m2.height; k++ {
				k1 := m1.Get(row, k)
				k2 := m2.Get(k, col)
				k3 := valsOp(k1, k2)
				toAgg[k] = k3 //valsOp(m1.Get(row, k), m2.Get(k, col))
			}

			result[row][col] = aggOp(toAgg...)
		}
	}

	return MatrixFromSlice(result)
}
