package tests

import (
	"testing"

	im "github.com/eightlay/InfiniteMath"
)

func TestAddScalar(t *testing.T) {
	m := im.NewMatrix([][]int{
		{1, 2},
		{3, 4},
	})
	r := im.NewMatrix([][]int{
		{2, 3},
		{4, 5},
	})

	m.AddScalar(1)

	if !m.Equal(r) {
		t.Fatalf("scalar addition error")
	}
}

func TestMultScalar(t *testing.T) {
	m := im.NewMatrix([][]int{
		{1, 2},
		{3, 4},
	})
	r := im.NewMatrix([][]int{
		{2, 4},
		{6, 8},
	})

	m.MultScalar(2)

	if !m.Equal(r) {
		t.Fatalf("scalar multiplication error")
	}
}
