package storage

import (
	"errors"
	"github.com/strwrd/driver-aood/src/domain"
	"reflect"
	"testing"
)

func TestStorage_InsertDriver(t *testing.T) {
	storage := &Storage{drivers: []domain.Driver{}}
	driver := &domain.Driver{ID: "driver-001"}

	if err := storage.InsertDriver(driver); err != nil {
		t.Fatalf("expected to be success insert driver, got %v", err)
	}

	if len(storage.drivers) != 1 {
		t.Fatalf("expected driver len %d, got %d", 1, len(storage.drivers))
	}

	if !reflect.DeepEqual(*driver, storage.drivers[0]) {
		t.Fatalf("expected driver %v, got %v", *driver, storage.drivers[0])
	}
}

func TestStorage_InsertDriverDuplicate(t *testing.T) {
	storage := &Storage{drivers: []domain.Driver{{ID: "driver-001"}}}
	driver := &domain.Driver{ID: "driver-001"}

	if err := storage.InsertDriver(driver); err != nil && !errors.Is(err, ErrDuplicateDriver) {
		t.Fatalf("expected to throw ErrDuplicateDriver, got %v", err)
	}
}
