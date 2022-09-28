package cmd

import (
	"fmt"
)

func (c *Command) RegisterDriver(name string) (string, error) {
	driver, err := c.ucase.RegisterDriver(name)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Driver %s registered", driver.ID), nil
}
