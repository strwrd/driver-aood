package storage

// DeleteDriverPool remove driver from driver pool
func (s *Storage) DeleteDriverPool(idx int) {
	// return when driver pool was empty
	if len(s.driverPool) == 0 {
		return
	}

	// delete on rear
	if idx == len(s.driverPool)-1 {
		s.driverPool = s.driverPool[:idx]
		return
	}

	s.driverPool = append(s.driverPool[:idx], s.driverPool[idx+1:]...)
}
