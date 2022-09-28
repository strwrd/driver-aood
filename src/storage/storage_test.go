package storage

import (
	"github.com/strwrd/driver-aood/src/domain"
	"os"
	"testing"
)

func TestNewStorageInit(t *testing.T) {
	// pre-setup
	os.Remove(FILE_PATH)

	storage, err := NewStorage()
	if err != nil {
		t.Fatalf("expected success, got %s\n", err.Error())
	}

	if len(storage.drivers) != 0 {
		t.Fatalf("expected len driver %d, got %d", 0, len(storage.drivers))
	}

	if len(storage.bookings) != 0 {
		t.Fatalf("expected len driver %d, got %d", 0, len(storage.bookings))
	}

	if len(storage.driverPool) != 0 {
		t.Fatalf("expected len driver %d, got %d", 0, len(storage.driverPool))
	}
}

func TestNewStorageExist(t *testing.T) {
	// pre-setup
	os.Remove(FILE_PATH)
	storage := &Storage{
		drivers:    []domain.Driver{{ID: "driver-001"}},
		bookings:   []domain.Booking{{ID: "booking-1", DriverID: "driver-001", IsCompleted: true, Distance: 15}},
		driverPool: []string{"driver-001"},
	}
	storage.Save()

	newStorage, err := NewStorage()
	if err != nil {
		t.Fatalf("expected success, got %s\n", err.Error())
	}

	if len(newStorage.drivers) != 1 {
		t.Fatalf("expected len driver %d, got %d", 1, len(newStorage.drivers))
	}

	if len(newStorage.bookings) != 1 {
		t.Fatalf("expected len driver %d, got %d", 1, len(newStorage.bookings))
	}

	if len(newStorage.driverPool) != 1 {
		t.Fatalf("expected len driver %d, got %d", 1, len(newStorage.driverPool))
	}
}

func TestStorage_Save(t *testing.T) {
	// pre-setup
	os.Remove(FILE_PATH)

	storage := &Storage{
		drivers:    []domain.Driver{{ID: "driver-001"}},
		bookings:   []domain.Booking{{ID: "booking-1", DriverID: "driver-001", IsCompleted: true, Distance: 15}},
		driverPool: []string{"driver-001"},
	}
	storage.Save()

	_, err := os.Open(FILE_PATH)
	if err != nil {
		t.Fatalf("expected file exist, got %s\n", err.Error())
	}

	// post-setup
	os.Remove(FILE_PATH)

}
