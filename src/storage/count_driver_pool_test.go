package storage

import (
	"github.com/strwrd/driver-aood/src/domain"
	"testing"
)

func TestStorage_CountDriverPool(t *testing.T) {
	storage := &Storage{
		drivers:    []domain.Driver{{ID: "driver-001"}},
		bookings:   []domain.Booking{{ID: "booking-1", DriverID: "driver-001", IsCompleted: true, Distance: 15}},
		driverPool: []string{"driver-001"},
	}

	if storage.CountDriverPool() != len(storage.driverPool) {
		t.Fatalf("expected booking len %d, got %d", len(storage.driverPool), storage.CountDriverPool())
	}
}
