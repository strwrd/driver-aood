package business

import (
	"errors"
	"testing"
)

func TestCase_CompleteBooking(t *testing.T) {
	strg := storageCreator(t)
	ucase := NewCase(strg)

	booking, err := ucase.CompleteBooking("booking-2")
	if err != nil {
		t.Fatalf("expected to be success complete booking, got %v", err)
	}

	if !booking.IsCompleted {
		t.Fatalf("expected booking is completed, got %t", booking.IsCompleted)
	}

	if strg.CountDriverPool() != 2 {
		t.Fatalf("expected driver pool len is %d, got %d", 2, strg.CountDriverPool())
	}
}

func TestCase_CompleteBookingAlreadyComplete(t *testing.T) {
	strg := storageCreator(t)
	ucase := NewCase(strg)

	if _, err := ucase.CompleteBooking("booking-1"); err != nil && !errors.Is(err, ErrBookingIsCompleted) {
		t.Fatalf("expected to throw ErrBookingIsCompleted, got %v", err)
	}
}

func TestCase_CompleteBookingNotFound(t *testing.T) {
	strg := storageCreator(t)
	ucase := NewCase(strg)

	if _, err := ucase.CompleteBooking("booking-3"); err != nil && !errors.Is(err, ErrBookingNotFound) {
		t.Fatalf("expected to throw ErrBookingNotFound, got %v", err)
	}
}
