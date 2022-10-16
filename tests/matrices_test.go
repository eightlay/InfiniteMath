package tests

import (
	"testing"

	im "github.com/eightlay/InfiniteMath"
)

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
