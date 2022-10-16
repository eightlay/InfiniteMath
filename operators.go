package infinitemath

import c "github.com/eightlay/InfiniteMath/iternal/constraints"

// Add operator
func addOp[T c.Numeric](vals ...T) T {
	var result T = 0
	for _, v := range vals {
		result += v
	}
	return result
}

// Mult operator
func multOp[T c.Numeric](vals ...T) T {
	var result T = 1
	for _, v := range vals {
		result *= v
	}
	return result
}
