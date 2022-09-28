package storage

import (
	"github.com/strwrd/driver-aood/src/domain"
	"reflect"
	"testing"
)

func TestStorage_GetBookingByID(t *testing.T) {
	tests := []struct {
		name  string
		input struct {
			storage *Storage
			id      string
		}
		expected struct {
			booking domain.Booking
			found   bool
		}
	}{
		{
			name: "search not found",
			input: struct {
				storage *Storage
				id      string
			}{
				storage: &Storage{bookings: []domain.Booking{
					{
						ID:          "booking-1",
						DriverID:    "driver-001",
						IsCompleted: false,
						Distance:    1,
					},
				}},
				id: "booking-2",
			},
			expected: struct {
				booking domain.Booking
				found   bool
			}{
				booking: domain.Booking{},
				found:   false,
			},
		},
		{
			name: "search found",
			input: struct {
				storage *Storage
				id      string
			}{
				storage: &Storage{bookings: []domain.Booking{
					{
						ID:          "booking-1",
						DriverID:    "driver-001",
						IsCompleted: false,
						Distance:    1,
					},
				}},
				id: "booking-1",
			},
			expected: struct {
				booking domain.Booking
				found   bool
			}{
				booking: domain.Booking{
					ID:          "booking-1",
					DriverID:    "driver-001",
					IsCompleted: false,
					Distance:    1,
				},
				found: true,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			booking, found := test.input.storage.GetBookingByID(test.input.id)

			if found != test.expected.found {
				t.Fatalf("expected found value %t, got %t", test.expected.found, found)
			}

			if !reflect.DeepEqual(booking, test.expected.booking) {
				t.Fatalf("expected booking value %v, got %v", test.expected.booking, booking)
			}
		})

	}
}
