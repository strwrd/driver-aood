package cmd

import "errors"

var (
	ErrInvalidDistanceInput = errors.New("error invalid distance input value")
	ErrInvalidMinimumInput  = errors.New("error invalid minimum input")
)
