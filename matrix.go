package infinitemath

import (
	c "github.com/eightlay/InfiniteMath/iternal/constraints"
	e "github.com/eightlay/InfiniteMath/iternal/errors"
)

// Matrix 2D of any numeric type
type Matrix[T c.Numeric] struct {
	vals   [][]T
	width  uint
	height uint
}

// Create new matrix from the given slice of shape NxM
func NewMatrix[T c.Numeric](vals [][]T) *Matrix[T] {
	height := uint(len(vals))
	if height == 0 {
		panic(e.ErrNullMatrix)
	}

	width := uint(len(vals[0]))
	if width == 0 {
		panic(e.ErrNullMatrix)
	}

	m := &Matrix[T]{
		vals:   make([][]T, height),
		width:  width,
		height: uint(height),
	}

	var row uint = 0
	for ; row < height-1; row++ {
		if len(vals[row]) != len(vals[row+1]) {
			panic(e.ErrRowSize)
		}

		m.vals[row] = make([]T, width)

		for col := uint(0); col < width; col++ {
			m.vals[row][col] = vals[row][col]
		}
	}

	m.vals[row] = make([]T, width)

	for col := uint(0); col < width; col++ {
		m.vals[row][col] = vals[row][col]
	}

	return m
}

// Get matrix shape
func (m *Matrix[T]) Shape() (uint, uint) {
	return m.height, m.width
}

// Get element
func (m *Matrix[T]) Get(row, col uint) T {
	if row >= m.height {
		panic(e.ErrRowIndex)
	}

	if col >= m.width {
		panic(e.ErrColIndex)
	}

	return m.vals[row][col]
}

// Get row
func (m *Matrix[T]) GetRow(row uint) *Matrix[T] {
	if row >= m.height {
		panic(e.ErrRowIndex)
	}

	return NewMatrix(m.vals[row : row+1])
}

// Get column
func (m *Matrix[T]) GetCol(col uint) *Matrix[T] {
	result := []T{}

	for row := uint(0); row < m.height; row++ {
		result = append(result, m.vals[row][col])
	}
}

// Corresponding matrices' elements are equal
func (m *Matrix[T]) Equal(rightMatrix *Matrix[T]) bool {
	if m.height != rightMatrix.height || m.width != rightMatrix.width {
		panic(e.ErrDimensions)
	}

	for row := uint(0); row < m.height; row++ {
		for col := uint(0); col < m.width; col++ {
			if m.vals[row][col] != rightMatrix.vals[row][col] {
				return false
			}
		}
	}

	return true
}

// Element-wise addition of matrix elements with
// elements of the rightMatrix
func (m *Matrix[T]) Add(rightMatrix *Matrix[T]) {
	Broadcast(m, rightMatrix, addOp[T])
}

// Element-wise multiplication of matrix elements
// by elements of the rightMatrix
func (m *Matrix[T]) Mult(rightMatrix *Matrix[T]) {
	Broadcast(m, rightMatrix, multOp[T])
}
