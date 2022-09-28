package business

import (
	"github.com/strwrd/driver-aood/src/domain"
	"github.com/strwrd/driver-aood/src/storage"
	"testing"
)

func storageCreator(t *testing.T) *storage.Storage {
	t.Helper()

	strg, err := storage.NewStorage()
	if err != nil {
		t.Fatalf("expected success to create storage, got %v", err)
	}

	if err := strg.InsertDriver(&domain.Driver{ID: "driver-001"}); err != nil {
		t.Fatalf("expected success to create driver, got %v", err)
	}

	if err := strg.InsertDriver(&domain.Driver{ID: "driver-002"}); err != nil {
		t.Fatalf("expected success to create driver, got %v", err)
	}

	booking1 := &domain.Booking{
		ID:          "booking-1",
		DriverID:    "driver-001",
		IsCompleted: true,
		Distance:    5,
	}

	booking2 := &domain.Booking{
		ID:          "booking-2",
		DriverID:    "driver-002",
		IsCompleted: false,
		Distance:    5,
	}

	strg.InsertBooking(booking1)
	strg.InsertBooking(booking2)

	if err := strg.InsertDriverPool("driver-001"); err != nil {
		t.Fatalf("expected success to insert driver pool, got %v", err)
	}

	return strg
}

func TestNewCase(t *testing.T) {
	storage := storageCreator(t)
	ucase := NewCase(storage)

	if ucase.Storage == nil {
		t.Fatalf("expected storage not nil, but %v", nil)
	}
}
