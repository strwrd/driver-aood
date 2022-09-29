package cmd

import (
	"fmt"
	"strconv"
)

func (c *Command) CompletedBookingGT(min string) (string, error) {
	value, err := strconv.Atoi(min)
	if err != nil {
		return "", ErrInvalidMinimumInput
	}

	completedBooking := c.ucase.CompletedBookingGT(value)

	output := "Driver Completed\n"

	for driverID, total := range completedBooking {
		output += fmt.Sprintf("%s %d\n", driverID, total)
	}

	return output, nil
}
