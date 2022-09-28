package cmd

import (
	"errors"
	"fmt"
	"github.com/strwrd/driver-aood/src/business"
	"strconv"
)

func (c *Command) DispatchDriver(distance string) (string, error) {
	dist, err := strconv.Atoi(distance)
	if err != nil {
		return "", ErrInvalidDistanceInput
	}

	booking, err := c.ucase.DispatchDriver(dist)
	if err != nil && !errors.Is(err, business.ErrDriverPoolEmpty) {
		return "", err
	}

	if errors.Is(err, business.ErrDriverPoolEmpty) {
		return fmt.Sprintf("Sorry, driver is not available"), nil
	}

	output := fmt.Sprintf(
		"Driver %s is assigned to booking %s with %d km distance",
		booking.DriverID,
		booking.ID,
		booking.Distance,
	)
	return output, nil
}
