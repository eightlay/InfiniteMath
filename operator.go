package infinitemath

import (
	c "github.com/eightlay/InfiniteMath/iternal/constraints"
)

// Operator over one or more matrix elements
type Operator[T c.Numeric] func(...T) T
