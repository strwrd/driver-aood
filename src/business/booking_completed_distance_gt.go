package business

func (c *Case) BookingCompletedDistanceGT(min int) map[string]int {
	drivers := c.Storage.GetDriver()
	result := make(map[string]int)

	for _, driver := range drivers {
		orders := c.Storage.GetBookingByDriverID(driver.ID)
		totalDist := 0

		for _, order := range orders {
			if order.IsCompleted {
				totalDist += order.Distance
			}
		}

		if totalDist > min {
			result[driver.ID] = totalDist
		}
	}

	return result
}
