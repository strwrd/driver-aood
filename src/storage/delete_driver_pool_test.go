package storage

import (
	"testing"
)

func TestStorage_DeleteDriverPool(t *testing.T) {
	tests := []struct {
		name  string
		input struct {
			storage *Storage
			idx     int
		}
		expected []string
	}{
		{
			name: "remove front",
			input: struct {
				storage *Storage
				idx     int
			}{
				storage: &Storage{driverPool: []string{"driver-001", "driver-002", "driver-003"}},
				idx:     0,
			},
			expected: []string{"driver-002", "driver-003"},
		},
		{
			name: "remove rear",
			input: struct {
				storage *Storage
				idx     int
			}{
				storage: &Storage{driverPool: []string{"driver-001", "driver-002", "driver-003"}},
				idx:     2,
			},
			expected: []string{"driver-001", "driver-002"},
		},
		{
			name: "remove mid",
			input: struct {
				storage *Storage
				idx     int
			}{
				storage: &Storage{driverPool: []string{"driver-001", "driver-002", "driver-003"}},
				idx:     1,
			},
			expected: []string{"driver-001", "driver-003"},
		},
		{
			name: "remove empty",
			input: struct {
				storage *Storage
				idx     int
			}{
				storage: &Storage{driverPool: []string{}},
				idx:     1,
			},
			expected: []string{},
		},
		{
			name: "remove last element",
			input: struct {
				storage *Storage
				idx     int
			}{
				storage: &Storage{driverPool: []string{"driver-001"}},
				idx:     0,
			},
			expected: []string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.input.storage.DeleteDriverPool(test.input.idx)

			if len(test.input.storage.driverPool) != len(test.expected) {
				t.Fatalf("expected driver pool len %d, got %d", len(test.expected), len(test.input.storage.driverPool))
			}

			for i := 0; i < len(test.expected); i++ {
				if test.input.storage.driverPool[i] != test.expected[i] {
					t.Fatalf("expected driver id %s, got %s", test.expected[i], test.input.storage.driverPool[i])
				}
			}
		})

	}
}
