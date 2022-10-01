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

	var i uint = 0
	for ; i < height-1; i++ {
		if len(vals[i]) != len(vals[i+1]) {
			panic(e.ErrRowSize)
		}

		m.vals[i] = make([]T, width)

		for j := uint(0); j < width; j++ {
			m.vals[i][j] = vals[i][j]
		}
	}

	m.vals[i] = make([]T, width)

	for j := uint(0); j < width; j++ {
		m.vals[i][j] = vals[i][j]
	}

	return m
}

// Corresponding matrices' elements are equal
func (m *Matrix[T]) Equal(rightMatrix *Matrix[T]) bool {
	if m.height != rightMatrix.height || m.width != rightMatrix.width {
		panic(e.ErrDimensions)
	}

	for i := uint(0); i < m.height; i++ {
		for j := uint(0); j < m.width; j++ {
			if m.vals[i][j] != rightMatrix.vals[i][j] {
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
