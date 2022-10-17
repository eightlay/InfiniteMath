package infinitemath

import (
	c "github.com/eightlay/InfiniteMath/iternal/constraints"
	e "github.com/eightlay/InfiniteMath/iternal/errors"
)

// Apply operator to each matrix element
func Apply[T c.Numeric](m *Matrix[T], f func(T) T) *Matrix[T] {
	result := m.Copy()
	result.Apply(f)
	return result
}

// Apply `by` function to `from` matrix and write result in `to` matrix
func WriteApply[T c.Numeric](to, from *Matrix[T], by func(T) T) {
	if to.height != from.height || to.width != from.width {
		panic(e.ErrWriteInDimensions)
	}

	for row := uint(0); row < to.height; row++ {
		for col := uint(0); col < to.width; col++ {
			to.vals[row][col] = by(from.vals[row][col])
		}
	}
}

// Apply operator along axis
func ApplyAlongAxis[Ti c.Numeric, To c.Numeric](m *Matrix[Ti], rowWise bool, f func([]Ti) To) *Matrix[To] {
	var getFunc func(row, col uint) []Ti
	var itRow, itCol uint

	if rowWise {
		getFunc = func(row, col uint) []Ti {
			return m.Row(row).AsSlice()[0]
		}
		itRow, itCol = m.height, 1
	} else {
		getFunc = func(row, col uint) []Ti {
			return m.Col(col).Transpose().AsSlice()[0]
		}
		itRow, itCol = 1, m.width
	}

	result := ZeroMatrix[To](itRow, itCol)

	for row := uint(0); row < itRow; row++ {
		for col := uint(0); col < itCol; col++ {
			result.vals[row][col] = f(getFunc(row, col))
		}
	}

	return result
}
