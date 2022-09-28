package storage

import "github.com/strwrd/driver-aood/src/domain"

// GetBookingByID get booking by booking id, return false if booking object not found
func (s *Storage) GetBookingByID(id string) (domain.Booking, bool) {
	for _, booking := range s.bookings {
		if booking.ID == id {
			return booking, true
		}
	}

	return domain.Booking{}, false
}
