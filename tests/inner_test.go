package tests

import (
	"testing"

	im "github.com/eightlay/InfiniteMath"
)

func TestDot(t *testing.T) {
	m := im.NewMatrix([][]int{
		{1, 2},
		{3, 4},
	})
	r := im.NewMatrix([][]int{
		{7, 10},
		{15, 22},
	})

	d := im.Dot(m, m)

	if !d.Equal(r) {
		t.Fatalf("dot product error")
	}
}
