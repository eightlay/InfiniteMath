package infinitemath

import (
	"math/rand"
	"time"

	c "github.com/eightlay/InfiniteMath/iternal/constraints"
	e "github.com/eightlay/InfiniteMath/iternal/errors"
)

// Generate matrix with given dimensions and generation function
func MatrixGenerator[T c.Numeric](height, width uint, genFunc func(height, width uint) T) *Matrix[T] {
	if height == 0 {
		panic(e.ErrNullMatrix)
	}

	if width == 0 {
		panic(e.ErrNullMatrix)
	}

	m := &Matrix[T]{
		vals:   make([][]T, height),
		width:  width,
		height: uint(height),
	}

	for row := uint(0); row < height-1; row++ {
		m.vals[row] = make([]T, width)

		for col := uint(0); col < width; col++ {
			m.vals[row][col] = genFunc(row, col)
		}
	}

	return m
}

// Generate height x width with the given value
func FillMatrix[T c.Numeric](height, width uint, val T) *Matrix[T] {
	valGen := func(height, width uint) T {
		return val
	}

	return MatrixGenerator(height, width, valGen)
}

// Generate height x width zero matrix
func ZeroMatrix[T c.Numeric](height, width uint) *Matrix[T] {
	return FillMatrix[T](height, width, 0)
}

// Generate height x height eye matrix
func EyeMatrix[T c.Numeric](height uint) *Matrix[T] {
	eye := func(height, width uint) T {
		if height == width {
			return 1
		}
		return 0
	}

	return MatrixGenerator(height, height, eye)
}

// Generate height x width matrix with random values in [a, b)
func RandomMatrix[T c.Numeric](height, width uint, a, b T) *Matrix[T] {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	if a > b {
		a, b = b, a
	}
	rng := float64(b - a)

	random := func(height, width uint) T {
		return T(r.Float64()*rng) + a
	}

	return MatrixGenerator(height, width, random)
}
