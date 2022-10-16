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

	im.Apply(m, f)

	if m.Get(0, 0) != 4 || m.Get(0, 1) != 9 {
		t.Fatal("invalid matrix creation")
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
