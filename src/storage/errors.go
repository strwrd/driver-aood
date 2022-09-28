package storage

import "errors"

var (
	ErrDuplicateDriver          = errors.New("error duplicate driver id")
	ErrBookingIdNotFound        = errors.New("error booking id not found")
	ErrDuplicateDriverPoolEntry = errors.New("error duplicate driver id on pool")
)
