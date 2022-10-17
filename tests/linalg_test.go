package tests

import (
	"testing"

	im "github.com/eightlay/InfiniteMath"
)

func TestDot(t *testing.T) {
	m := im.MatrixFromSlice([][]int{
		{1, 2},
		{3, 4},
	})
	r := im.MatrixFromSlice([][]int{
		{7, 10},
		{15, 22},
	})

	d := im.Dot(m, m)

	if !d.Equal(r) {
		t.Fatalf("dot product error")
	}
}

func TestWriteDot(t *testing.T) {
	m := im.ZeroMatrix[int](2, 2)
	m1 := im.FillMatrix(1, 2, 2)
	m2 := im.FillMatrix(2, 1, 3)

	im.WriteDot(m, m1, m2)

	if !(m.Get(0, 0) == 12) {
		t.Fatalf("dot product error")
	}
}
