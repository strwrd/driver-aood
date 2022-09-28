package business

import "github.com/strwrd/driver-aood/src/domain"

func (c *Case) CompleteBooking(bookingID string) (*domain.Booking, error) {
	// get and validate booking
	booking, found := c.Storage.GetBookingByID(bookingID)
	if !found {
		return nil, ErrBookingNotFound
	}

	if booking.IsCompleted {
		return nil, ErrBookingIsCompleted
	}

	// update booking property
	booking.IsCompleted = true
	if err := c.Storage.UpdateBooking(&booking); err != nil {
		return nil, err
	}

	// return driver into driver pool
	c.Storage.InsertDriverPool(booking.DriverID)
	return &booking, nil
}
