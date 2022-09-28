package storage

import "github.com/strwrd/driver-aood/src/domain"

// InsertBooking insert booking object
func (s *Storage) InsertBooking(b *domain.Booking) {
	s.bookings = append(s.bookings, *b)
}
