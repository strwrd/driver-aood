package storage

// CountBooking return total count of booking
func (s *Storage) CountBooking() int {
	return len(s.bookings)
}
