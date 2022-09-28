package storage

import (
	"errors"
	"testing"
)

func TestStorage_InsertDriverPool(t *testing.T) {
	storage := &Storage{driverPool: []string{}}
	driverID := "driver-001"

	if err := storage.InsertDriverPool(driverID); err != nil {
		t.Fatalf("expected to be success insert driverID into driver pool, got %v", err)
	}

	if len(storage.driverPool) != 1 {
		t.Fatalf("expected driver pool len %d, got %d", 1, len(storage.driverPool))
	}

	if storage.driverPool[0] != driverID {
		t.Fatalf("expected driver id %s, got %s", driverID, storage.driverPool[0])
	}
}

func TestStorage_InsertDriverPoolDuplicate(t *testing.T) {
	storage := &Storage{driverPool: []string{"driver-001"}}
	driverID := "driver-001"

	if err := storage.InsertDriverPool(driverID); err != nil && !errors.Is(err, ErrDuplicateDriverPoolEntry) {
		t.Fatalf("expected to throw ErrDuplicateDriverPoolEntry, got %v", err)
	}
}
