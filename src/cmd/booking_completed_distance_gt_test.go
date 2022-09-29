package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

func TestCommand_BookingCompletedDistanceGT(t *testing.T) {
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

	result, err := command.BookingCompletedDistanceGT("12")
	if err != nil {
		t.Fatalf("expected to be success execute command BookingCompletedDistanceGT, got %v", err)
	}

	expResult := "Driver Completed\ndriver-001 17\n"

	if result != expResult {
		t.Fatalf("expected output `%s`, got `%s`", expResult, result)
	}
}

func TestCommand_BookingCompletedDistanceGTInvalid(t *testing.T) {
	ucase := businessCreator(t)
	command := NewCommand(ucase)

	_, err := command.BookingCompletedDistanceGT("a")
	if err != nil && !errors.Is(err, ErrInvalidMinimumInput) {
		t.Fatalf("expected to be throw ErrInvalidMinimumInput, got %v", err)
	}
}
