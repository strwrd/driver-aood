package storage

import (
	"github.com/strwrd/driver-aood/src/domain"
	"testing"
)

func TestStorage_GetBookingByDriverID(t *testing.T) {
	storage := &Storage{bookings: []domain.Booking{
		{
			ID:          "booking-1",
			DriverID:    "driver-001",
			IsCompleted: false,
			Distance:    0,
		},
		{
			ID:          "booking-2",
			DriverID:    "driver-001",
			IsCompleted: false,
			Distance:    0,
		},
		{
			ID:          "booking-3",
			DriverID:    "driver-002",
			IsCompleted: false,
			Distance:    0,
		},
	}}

	bookings := storage.GetBookingByDriverID("driver-001")

	if len(bookings) != 2 {
		t.Fatalf("expected len booking %d, got %d", 2, len(bookings))
	}
}
