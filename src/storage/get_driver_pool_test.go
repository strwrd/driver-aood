package storage

import "testing"

func TestStorage_GetDriverPool(t *testing.T) {
	storage := &Storage{driverPool: []string{"driver-001", "driver-002", "driver-003"}}
	result := storage.GetDriverPool()

	if len(result) != len(storage.driverPool) {
		t.Fatalf("expected len driver pool %d, got %d", len(storage.driverPool), len(result))
	}

	for i := 0; i < len(storage.driverPool); i++ {
		if result[i] != storage.driverPool[i] {
			t.Fatalf("expected driver id %s, got %s", result[i], storage.driverPool[i])
		}
	}
}
