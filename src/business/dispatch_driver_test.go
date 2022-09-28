package business

import (
	"errors"
	"github.com/strwrd/driver-aood/src/domain"
	"reflect"
	"testing"
)

func TestCase_DispatchDriver(t *testing.T) {
	strg := storageCreator(t)
	ucase := NewCase(strg)

	booking, err := ucase.DispatchDriver(100)
	if err != nil {
		t.Fatalf("expected to be success dispatch driver, got %v", err)
	}

	expBooking := &domain.Booking{
		ID:          "booking-3",
		DriverID:    "driver-001",
		IsCompleted: false,
		Distance:    100,
	}

	if !reflect.DeepEqual(booking, expBooking) {
		t.Fatalf("expected booking is %v, got %v", expBooking, booking)
	}

	if strg.CountBooking() != 3 {
		t.Fatalf("expected booking len is %d, got %d", 3, strg.CountBooking())
	}

	if strg.CountDriverPool() != 0 {
		t.Fatalf("expected driver pool len is %d, got %d", 0, strg.CountDriverPool())
	}
}

func TestCase_DispatchDriverEmptyPool(t *testing.T) {
	strg := storageCreator(t)
	ucase := NewCase(strg)
	strg.DeleteDriverPool(0)

	_, err := ucase.DispatchDriver(100)
	if err != nil && !errors.Is(err, ErrDriverPoolEmpty) {
		t.Fatalf("expected to throw ErrDriverPoolEmpty, got %v", err)
	}
}
