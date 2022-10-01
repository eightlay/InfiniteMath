package infinitemath

import c "github.com/eightlay/InfiniteMath/iternal/constraints"

func addOp[T c.Numeric](v1 T, v2 T) T {
	return v1 + v2
}

func multOp[T c.Numeric](v1 T, v2 T) T {
	return v1 * v2
}
