package infinitemath

import c "github.com/eightlay/InfiniteMath/iternal/constraints"

func addOp[T c.Numeric](vals ...T) T {
	var result T = 0
	for _, v := range vals {
		result += v
	}
	return result
}

func multOp[T c.Numeric](vals ...T) T {
	var result T = 1
	for _, v := range vals {
		result *= v
	}
	return result
}
