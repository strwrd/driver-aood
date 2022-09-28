package storage

import "github.com/strwrd/driver-aood/src/domain"

// UpdateBooking  update booking object
func (s *Storage) UpdateBooking(booking *domain.Booking) error {
	// check if booking object exist
	idx := -1
	for i, book := range s.bookings {
		if book.ID == booking.ID {
			idx = i
			break
		}
	}
	if idx == -1 {
		return ErrBookingIdNotFound
	}

	// update booking object of current pointer (replace)
	s.bookings[idx] = *booking

	return nil
}
