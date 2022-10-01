package tests

import (
	"testing"

	im "github.com/eightlay/InfiniteMath"
)

func eye1Matrices() (*im.Matrix[int], *im.Matrix[int], *im.Matrix[int]) {
	m1 := im.NewMatrix([][]int{
		{1, 0},
		{0, 1},
	})
	m2 := im.NewMatrix([][]int{
		{1, 0},
		{0, 1},
	})
	msum := im.NewMatrix([][]int{
		{2, 0},
		{0, 2},
	})
	return m1, m2, msum
}

func TestAdd(t *testing.T) {
	m1, m2, msum := eye1Matrices()

	m1.Add(m2)

	if !m1.Equal(msum) {
		t.Fatalf("addition error")
	}
}

func TestMult(t *testing.T) {
	m1, m2, _ := eye1Matrices()

	m1.Mult(m2)

	if !m1.Equal(m1) {
		t.Fatalf("multiplication error")
	}
}
