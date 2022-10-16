package infinitemath

import (
	c "github.com/eightlay/InfiniteMath/iternal/constraints"
)

// Operator over one or more matrix elements
type Operator[Ti c.Numeric, To c.Numeric] func(...Ti) To
