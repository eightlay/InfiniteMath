package tests

import (
	"testing"

	im "github.com/eightlay/InfiniteMath"
)

func TestMatrixFromSlice(t *testing.T) {
	m := im.MatrixFromSlice([][]int{{1, 2}})

	if m.Get(0, 0) != 1 || m.Get(0, 1) != 2 {
		t.Fatal("invalid matrix creation")
	}
}

func TestZeroMatrix(t *testing.T) {
	m := im.ZeroMatrix[int](2, 2)

	c := m.Get(0, 0) == 0
	c = c || m.Get(0, 1) == 0
	c = c || m.Get(1, 1) == 0
	c = c || m.Get(1, 0) == 0

	if !c {
		t.Fatal("invalid zero matrix creation")
	}
}

func TestEyeMatrix(t *testing.T) {
	m := im.ZeroMatrix[int](2, 2)

	c := m.Get(0, 0) == 1
	c = c || m.Get(0, 1) == 0
	c = c || m.Get(1, 1) == 1
	c = c || m.Get(1, 0) == 0

	if !c {
		t.Fatal("invalid eye matrix creation")
	}
}
