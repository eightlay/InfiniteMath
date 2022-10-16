package infinitemath

import (
	c "github.com/eightlay/InfiniteMath/iternal/constraints"
)

// Spread the operator with the given last argument to matrix elements
func Spread[T c.Numeric](m *Matrix[T], v T, op Operator[T]) {
	for row := uint(0); row < m.height; row++ {
		for col := uint(0); col < m.width; col++ {
			m.vals[row][col] = op(m.vals[row][col], v)
		}
	}
}
