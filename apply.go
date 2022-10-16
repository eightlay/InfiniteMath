package infinitemath

import (
	c "github.com/eightlay/InfiniteMath/iternal/constraints"
)

// Apply operator to each matrix element
func Apply[T c.Numeric](m *Matrix[T], f func(T) T) {
	for row := uint(0); row < m.height; row++ {
		for col := uint(0); col < m.width; col++ {
			m.vals[row][col] = f(m.vals[row][col])
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
