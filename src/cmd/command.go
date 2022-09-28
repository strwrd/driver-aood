package cmd

import (
	"github.com/strwrd/driver-aood/src/business"
)

type Command struct {
	ucase *business.Case
}

func NewCommand(ucase *business.Case) *Command {
	return &Command{ucase: ucase}
}
