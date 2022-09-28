package storage

import (
	"github.com/strwrd/driver-aood/src/domain"
	"reflect"
	"testing"
)

func TestStorage_InsertBooking(t *testing.T) {
	storage := &Storage{bookings: []domain.Booking{}}
	booking := &domain.Booking{
		ID:          "booking-1",
		DriverID:    "driver-001",
		IsCompleted: false,
		Distance:    1,
	}

	storage.InsertBooking(booking)

	if len(storage.bookings) != 1 {
		t.Fatalf("expected booking len %d, got %d", 1, len(storage.bookings))
	}

	if !reflect.DeepEqual(*booking, storage.bookings[0]) {
		t.Fatalf("expected booking %v, got %v", *booking, storage.bookings[0])
	}
}
