package tests

import (
	"testing"

	im "github.com/eightlay/InfiniteMath"
)

func TestCreation(t *testing.T) {
	m := im.MatrixFromSlice([][]int{{1, 2}})

	if m.Get(0, 0) != 1 || m.Get(0, 1) != 2 {
		t.Fatal("invalid matrix creation")
	}
}

func TestString(t *testing.T) {
	m := im.MatrixFromSlice([][]int{{1, 2}})
	desired := "Matrix[int]{{{1, 2}}}"

	if m.String() != desired {
		t.Fatal("invalid matrix creation")
	}
}

func TestCopy(t *testing.T) {
	m := im.MatrixFromSlice([][]int{{1, 2}})
	mc := m.Copy()

	if !m.Equal(mc) {
		t.Fatal("invalid matrix creation")
	}
}

func TestAsSlice(t *testing.T) {
	m := im.MatrixFromSlice([][]int{{1}, {2}})
	s := m.AsSlice()

	t1 := m.Get(0, 0) == s[0][0]
	t2 := m.Get(1, 0) == s[1][0]

	if !(t1 && t2) {
		t.Fatal("invalid matrix creation")
	}
}

func TestTranspose(t *testing.T) {
	m := im.MatrixFromSlice([][]int{{1, 2}})
	mc := im.MatrixFromSlice([][]int{{1}, {2}})
	mt := m.Transpose()

	if !mc.Equal(mt) {
		t.Fatal("invalid matrix creation")
	}
}
