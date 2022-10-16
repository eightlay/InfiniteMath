package infinitemath

import (
	c "github.com/eightlay/InfiniteMath/iternal/constraints"
)

// Apply operator to each matrix element
func Apply[T c.Numeric](m *Matrix[T], op Operator[T, T]) {
	for row := uint(0); row < m.height; row++ {
		for col := uint(0); col < m.width; col++ {
			m.vals[row][col] = op(m.vals[row][col])
		}
	}
}
