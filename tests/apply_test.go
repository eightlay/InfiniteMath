package tests

import (
	"testing"

	im "github.com/eightlay/InfiniteMath"
)

func TestApply(t *testing.T) {
	m := im.MatrixFromSlice([][]int{{2, 3}})

	f := func(v int) int {
		return v * v
	}

	m.Apply(f)

	if m.Get(0, 0) != 4 || m.Get(0, 1) != 9 {
		t.Fatal("invalid apply")
	}
}

func TestWriteApply(t *testing.T) {
	mto := im.ZeroMatrix[int](1, 1)
	mfrom := im.FillMatrix(1, 1, 2)

	f := func(v int) int {
		return v * v
	}

	im.WriteApply(mto, mfrom, f)

	if mto.Get(0, 0) != 4 {
		t.Fatal("invalid apply")
	}
}

func TestApplyAlongAxis(t *testing.T) {
	m := im.MatrixFromSlice([][]int{{2, 3}, {4, 5}})

	f := func(vals []int) int {
		s := 0

		for _, v := range vals {
			s += v
		}

		return s
	}

	n := im.ApplyAlongAxis(m, true, f)

	if n.Get(0, 0) != 5 || n.Get(1, 0) != 9 {
		t.Fatal("invalid matrix creation")
	}
}
