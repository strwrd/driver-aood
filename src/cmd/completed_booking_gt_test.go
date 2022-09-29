package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

func TestCommand_CompletedBookingGT(t *testing.T) {
	ucase := businessCreator(t)
	command := NewCommand(ucase)

	for i := 3; i < 6; i++ {
		_, err := command.DispatchDriver(strconv.Itoa(i))
		if err != nil {
			t.Fatalf("expected to be success dispatch driver, got %v", err)
		}

		_, err = command.CompleteBooking(fmt.Sprintf("booking-%d", i))
		if err != nil {
			t.Fatalf("expected to be success complete booking, got %v", err)
		}
	}

	result, err := command.CompletedBookingGT("3")
	if err != nil {
		t.Fatalf("expected to be success execute command completedBookingGT, got %v", err)
	}

	expResult := "Driver Completed\ndriver-001 4\n"

	if result != expResult {
		t.Fatalf("expected output `%s`, got `%s`", expResult, result)
	}
}

func TestCommand_CompletedBookingGTInvalid(t *testing.T) {
	ucase := businessCreator(t)
	command := NewCommand(ucase)

	_, err := command.CompletedBookingGT("a")
	if err != nil && !errors.Is(err, ErrInvalidMinimumInput) {
		t.Fatalf("expected to be throw ErrInvalidMinimumInput, got %v", err)
	}
}
