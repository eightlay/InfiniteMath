package infinitemath

import (
	"fmt"
	"math"
	"reflect"
	"strings"

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
func MatrixFromSlice[T c.Numeric](vals [][]T) *Matrix[T] {
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

// Copy matrix
func (m *Matrix[T]) Copy() *Matrix[T] {
	return MatrixFromSlice(m.vals)
}

// Get matrix as 2D slice
func (m *Matrix[T]) AsSlice() [][]T {
	result := make([][]T, m.height)

	for row := uint(0); row < m.height; row++ {
		result[row] = make([]T, m.width)
		copy(result[row], m.vals[row])
	}

	return result
}

// Get matrix shape
func (m *Matrix[T]) Shape() (uint, uint) {
	return m.height, m.width
}

// Transpose matrix
func (m *Matrix[T]) Transpose() *Matrix[T] {
	result := ZeroMatrix[T](m.width, m.height)

	for row := uint(0); row < m.height; row++ {
		for col := uint(0); col < m.width; col++ {
			result.vals[col][row] = m.vals[row][col]
		}
	}

	return result
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
func (m *Matrix[T]) Row(row uint) *Matrix[T] {
	if row >= m.height {
		panic(e.ErrRowIndex)
	}

	return MatrixFromSlice(m.vals[row : row+1])
}

// Get column
func (m *Matrix[T]) Col(col uint) *Matrix[T] {
	result := make([][]T, m.height)
	result[0] = make([]T, 1)

	for row := uint(0); row < m.height; row++ {
		result[row][0] = m.vals[row][col]
	}

	return MatrixFromSlice(result)
}

// Corresponding matrices' elements are equal
func (m *Matrix[T]) Equal(rightMatrix *Matrix[T]) bool {
	if m.height != rightMatrix.height || m.width != rightMatrix.width {
		panic(e.ErrBroadcastDimensions)
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

// Add scalar to each matrix element
func (m *Matrix[T]) AddScalar(s T) {
	Spread(m, s, addOp[T])
}

// Multiply each mutrix element by scalar
func (m *Matrix[T]) MultScalar(s T) {
	Spread(m, s, multOp[T])
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

// Convert matrix to its string represention
func (m *Matrix[T]) String() string {
	// Convert numbers to strings and find length of the longest one
	strs := make([][]string, m.height)
	maxLen := 0

	for row := uint(0); row < m.height; row++ {
		strs[row] = make([]string, m.width)

		for col := uint(0); col < m.width; col++ {
			strs[row][col] = fmt.Sprintf("%v", m.vals[row][col])

			l := len(strs[row][col])
			if l > maxLen {
				maxLen = l
			}
		}
	}

	// Create string represenation of the matrix
	result := fmt.Sprintf("Matrix[%v]{{", reflect.TypeOf((*T)(nil)).Elem().Name())
	rowPrefix := strings.Repeat(" ", len(result))
	offsetRow := 0
	offsetCol := 0

	for row := 0; row < int(math.Min(6, float64(m.height))); row++ {
		if row != 0 {
			result += ",\n" + rowPrefix
		}

		if row == 3 {
			result += "...,\n" + rowPrefix
			offsetRow = int(math.Max(float64(m.height)-6, 0))
		}

		result += "{"

		for col := 0; col < int(math.Min(6, float64(m.width))); col++ {
			if col == 3 {
				result += ", ..."
				offsetCol = int(math.Max(float64(m.width)-6, 0))
			}

			r, c := row+offsetRow, col+offsetCol

			prefix := strings.Repeat(" ", maxLen-len(strs[r][c]))

			if col != 0 {
				result += ", "
			}
			result += prefix + strs[r][c]
		}

		result += "}"
	}

	result += "}}"

	return result
}
