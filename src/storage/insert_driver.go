package storage

import (
	"fmt"
	"github.com/strwrd/driver-aood/src/domain"
)

// InsertDriver insert driver object
func (s *Storage) InsertDriver(d *domain.Driver) error {
	// validate if there is no conflict on driver id
	for _, driver := range s.drivers {
		if d.ID == driver.ID {
			return fmt.Errorf("duplicate entry: %w", ErrDuplicateDriver)
		}
	}

	s.drivers = append(s.drivers, *d)
	return nil
}
