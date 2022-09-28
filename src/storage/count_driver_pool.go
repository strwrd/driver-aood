package storage

// CountDriverPool return total count of available drivers
func (s *Storage) CountDriverPool() int {
	return len(s.driverPool)
}
