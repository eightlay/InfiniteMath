package main

import (
	"fmt"

	im "github.com/eightlay/InfiniteMath"
)

func main() {
	// Create matrix
	m1 := im.MatrixFromSlice([][]int{
		{1, 2},
		{3, 4},
	})
	m2 := im.MatrixFromSlice([][]int{
		{1, 0},
		{0, 1},
	})

	// Add scalar
	m1.AddScalar(1)
	fmt.Printf("%s\n", m1)

	// Multiply by scalar
	m2.MultScalar(2)
	fmt.Printf("%s\n", m2)

	// Element-wise addition
	m1.Add(m2)
	fmt.Printf("%s\n", m1)

	// Element-wise multiplication
	m2.Mult(m1)
	fmt.Printf("%s\n", m2)

	// Dot product
	m3 := im.Dot(m1, m2)
	fmt.Printf("%s\n", m3)
}
