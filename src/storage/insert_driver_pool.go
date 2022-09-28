package storage

// InsertDriverPool insert driver into driver pool
func (s *Storage) InsertDriverPool(id string) error {
	// check duplication driver id
	for _, driverID := range s.driverPool {
		if driverID == id {
			return ErrDuplicateDriverPoolEntry
		}
	}

	s.driverPool = append(s.driverPool, id)
	return nil
}
