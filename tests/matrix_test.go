package tests

import (
	"testing"

	im "github.com/eightlay/InfiniteMath"
)

func TestCreation(t *testing.T) {
	m := im.NewMatrix([][]int{{1, 2}})

	if m.Get(0, 0) != 1 || m.Get(0, 1) != 2 {
		t.Fatal("invalid matrix creation")
	}
}

func TestString(t *testing.T) {
	m := im.NewMatrix([][]int{{1, 2}})
	desired := "Matrix[int]{{{1, 2}}}"

	if m.String() != desired {
		t.Fatal("invalid matrix creation")
	}
}
