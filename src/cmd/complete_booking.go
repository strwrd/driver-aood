package cmd

import (
	"fmt"
)

func (c *Command) CompleteBooking(bookingID string) (string, error) {
	booking, err := c.ucase.CompleteBooking(bookingID)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Driver %s is released to allocation pool", booking.DriverID), nil
}
