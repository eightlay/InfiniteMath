package errors

import "fmt"

var (
	ErrBroadcastDimensions = fmt.Errorf("matrices must have the same shape")
	ErrWriteInDimensions   = fmt.Errorf("matrices must have the same shape")
	ErrInnerDimensions     = fmt.Errorf("the left matrix width must be equal to the right matrix height")
	ErrNullMatrix          = fmt.Errorf("matrix can't be empty")
	ErrRowSize             = fmt.Errorf("length of matrix rows must be the same")
	ErrRowIndex            = fmt.Errorf("row index is out of range")
	ErrColIndex            = fmt.Errorf("column index is out of range")
)
