package storage

import (
	"github.com/strwrd/driver-aood/src/domain"
	"testing"
)

func TestStorage_UpdateBooking(t *testing.T) {
	booking := domain.Booking{
		ID:          "booking-1",
		DriverID:    "driver-001",
		IsCompleted: false,
		Distance:    10,
	}

	storage := &Storage{bookings: []domain.Booking{booking}}

	booking.IsCompleted = true
	storage.UpdateBooking(&booking)

	if !storage.bookings[0].IsCompleted {
		t.Fatalf("expected booking isCompleted true, got %t", storage.bookings[0].IsCompleted)
	}
}
