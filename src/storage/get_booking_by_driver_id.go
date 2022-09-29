package storage

import "github.com/strwrd/driver-aood/src/domain"

func (s *Storage) GetBookingByDriverID(driverID string) []domain.Booking {
	bookings := make([]domain.Booking, 0)

	for _, booking := range s.bookings {
		if booking.DriverID == driverID {
			bookings = append(bookings, booking)
		}
	}

	return bookings
}
