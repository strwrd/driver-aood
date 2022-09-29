package business

import (
	"fmt"
	"testing"
)

func TestCase_CompletedBookingGT(t *testing.T) {
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

	result := ucase.CompletedBookingGT(3)

	if len(result) != 1 {
		t.Fatalf("expected map len is %d, got %d", 1, len(result))
	}

	if result["driver-001"] != 4 {
		t.Fatalf("expected driver-001 value is %d, got %d", 4, result["driver-001"])
	}
}
