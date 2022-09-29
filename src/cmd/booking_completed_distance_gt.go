package cmd

import (
	"fmt"
	"strconv"
)

func (c *Command) BookingCompletedDistanceGT(distance string) (string, error) {
	value, err := strconv.Atoi(distance)
	if err != nil {
		return "", nil
	}

	result := c.ucase.BookingCompletedDistanceGT(value)

	output := "Driver Completed\n"
	for driverID, totalDist := range result {
		output += fmt.Sprintf("%s %d\n", driverID, totalDist)
	}

	return output, nil
}
