package business

import (
	"fmt"
	"github.com/strwrd/driver-aood/src/domain"
	"math/rand"
	"time"
)

func (c *Case) DispatchDriver(distance int) (*domain.Booking, error) {
	// check if current drivers are busy
	poolSize := c.Storage.CountDriverPool()
	if poolSize <= 0 {
		return nil, ErrDriverPoolEmpty
	}

	// randomize driver selection
	rand.Seed(time.Now().UnixNano())
	driverPool := c.Storage.GetDriverPool()

	// guard rand.Intn panic when n arg is <= 0
	poolIdx := 0
	if poolSize != 1 {
		poolIdx = rand.Intn(poolSize - 1)
	}

	driverID := driverPool[poolIdx]

	// create booking
	booking := &domain.Booking{
		ID:          fmt.Sprintf("booking-%d", c.Storage.CountBooking()+1),
		DriverID:    driverID, // pseudo random select driver
		IsCompleted: false,
		Distance:    distance,
	}
	c.Storage.InsertBooking(booking)

	// remove selected driver from driver pool
	c.Storage.DeleteDriverPool(poolIdx)

	return booking, nil
}
