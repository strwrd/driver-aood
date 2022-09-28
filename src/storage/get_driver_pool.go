package storage

// GetDriverPool return list of driver pool
func (s *Storage) GetDriverPool() []string {
	pool := make([]string, len(s.driverPool))
	copy(pool, s.driverPool)

	return pool
}
