package business

func (c *Case) CompletedBookingGT(min int) map[string]int {
	drivers := c.Storage.GetDriver()
	result := make(map[string]int)

	for _, driver := range drivers {
		orders := c.Storage.GetBookingByDriverID(driver.ID)
		counter := 0

		for _, order := range orders {
			if order.IsCompleted {
				counter += 1
			}
		}

		if counter >= min {
			result[driver.ID] = counter
		}
	}

	return result
}
