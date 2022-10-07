package infinitemath

import (
	c "github.com/eightlay/InfiniteMath/iternal/constraints"
)

type Operator[T c.Numeric] func(...T) T
