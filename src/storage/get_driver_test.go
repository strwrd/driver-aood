package storage

import (
	"github.com/strwrd/driver-aood/src/domain"
	"testing"
)

func TestStorage_GetDriver(t *testing.T) {
	storage := &Storage{drivers: []domain.Driver{
		{ID: "driver-001"},
		{ID: "driver-002"},
		{ID: "driver-003"},
	}}

	drivers := storage.GetDriver()

	if len(drivers) != 3 {
		t.Fatalf("expected driver len %d, got %d", 3, len(drivers))
	}
}
