package business

import "errors"

var (
	ErrDriverPoolEmpty    = errors.New("error driver pool is empty")
	ErrBookingNotFound    = errors.New("error booking id not found")
	ErrBookingIsCompleted = errors.New("error booking is already completed")
)
