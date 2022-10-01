package errors

import "fmt"

var (
	ErrDimensions = fmt.Errorf("operands could not be broadcast together")
	ErrNullMatrix = fmt.Errorf("matrix can't be empty")
	ErrRowSize    = fmt.Errorf("length of matrix rows must be the same")
)
