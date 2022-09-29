package business

import (
	"fmt"
	"testing"
)

func TestCase_BookingCompletedDistanceGT(t *testing.T) {
	strg := storageCreator(t)
	ucase := NewCase(strg)

	for i := 3; i < 6; i++ {
		_, err := ucase.DispatchDriver(i)
		if err != nil {
			t.Fatalf("expected to be success dispatch driver, got %v", err)
		}

		_, err = ucase.CompleteBooking(fmt.Sprintf("booking-%d", i))
		if err != nil {
			t.Fatalf("expected to be success complete booking, got %v", err)
		}
	}

	result := ucase.BookingCompletedDistanceGT(12)

	if len(result) != 1 {
		t.Fatalf("expected map len is %d, got %d", 1, len(result))
	}

	if result["driver-001"] != 17 {
		t.Fatalf("expected driver-001 value is %d, got %d", 17, result["driver-001"])
	}
}
