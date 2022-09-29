package storage

import "github.com/strwrd/driver-aood/src/domain"

func (s *Storage) GetDriver() []domain.Driver {
	drivers := make([]domain.Driver, len(s.drivers))
	copy(drivers, s.drivers)

	return drivers
}