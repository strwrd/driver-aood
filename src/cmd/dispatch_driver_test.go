package cmd

import (
	"errors"
	"testing"
)

func TestCommand_DispatchDriver(t *testing.T) {
	ucase := businessCreator(t)
	command := NewCommand(ucase)

	output, err := command.DispatchDriver("100")
	if err != nil {
		t.Fatalf("expected to be success dispatch driver, got %v", err)
	}

	expOutput := "Driver driver-001 is assigned to booking booking-3 with 100 km distance"

	if output != expOutput {
		t.Fatalf("expected output `%s`, got `%s`", expOutput, output)
	}
}

func TestCommand_DispatchDriverBadInput(t *testing.T) {
	ucase := businessCreator(t)
	command := NewCommand(ucase)

	if _, err := command.DispatchDriver("100q"); err != nil && !errors.Is(err, ErrInvalidDistanceInput) {
		t.Fatalf("expected to throw ErrInvalidDistanceInput, got %v", err)
	}
}

func TestCommand_DispatchDriverInvalid(t *testing.T) {
	ucase := businessCreator(t)
	command := NewCommand(ucase)

	if _, err := command.DispatchDriver("100"); err != nil {
		t.Fatalf("expected to be success dispatch driver, got %v", err)
	}

	output, err := command.DispatchDriver("50")
	if err != nil {
		t.Fatalf("expected to not throw error, got %v", err)
	}

	expOutput := "Sorry, driver is not available"

	if output != expOutput {
		t.Fatalf("expected output `%s`, got `%s`", expOutput, output)
	}
}
