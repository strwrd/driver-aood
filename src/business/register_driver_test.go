package business

import (
	"errors"
	"github.com/strwrd/driver-aood/src/storage"
	"testing"
)

func TestCase_RegisterDriver(t *testing.T) {
	strg := storageCreator(t)
	ucase := NewCase(strg)

	if _, err := ucase.RegisterDriver("driver-003"); err != nil {
		t.Fatalf("expected success to insert driver, got %v", err)
	}

	if strg.CountDriverPool() != 2 {
		t.Fatalf("expected driver pool len is %d, got %d", 2, strg.CountDriverPool())
	}
}

func TestCase_RegisterDriverDuplicate(t *testing.T) {
	strg := storageCreator(t)
	ucase := NewCase(strg)

	_, err := ucase.RegisterDriver("driver-001")
	if err != nil && !errors.Is(err, storage.ErrDuplicateDriver) {
		t.Fatalf("expected to throw storage.ErrDuplicateDriver, got %v", err)
	}
}
