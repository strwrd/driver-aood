package cmd

import (
	"testing"
)

func TestCommand_CompleteBooking(t *testing.T) {
	ucase := businessCreator(t)
	command := NewCommand(ucase)

	output, err := command.CompleteBooking("booking-2")
	if err != nil {
		t.Fatalf("expected to be able complete booking, got %v", err)
	}

	expOutput := "Driver driver-002 is released to allocation pool"

	if output != expOutput {
		t.Fatalf("expected output `%s`, got `%s`", expOutput, output)
	}
}

func TestCommand_CompleteBookingInvalid(t *testing.T) {
	ucase := businessCreator(t)
	command := NewCommand(ucase)

	if _, err := command.CompleteBooking("booking-1"); err == nil {
		t.Fatal("expected to throw error, got nil")
	}
}
