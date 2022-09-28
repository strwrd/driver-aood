package storage

import (
	"encoding/json"
	"errors"
	"github.com/strwrd/driver-aood/src/domain"
	"io"
	"os"
)

const (
	FILE_PATH = "./storage.json"
)

type Storage struct {
	drivers    []domain.Driver  `json:"drivers"`
	bookings   []domain.Booking `json:"orders"`
	driverPool []string         `json:"driverPool"`
}

func NewStorage() (*Storage, error) {
	file, err := os.Open(FILE_PATH)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	// initialize when file is not created
	if errors.Is(err, os.ErrNotExist) {
		storage := &Storage{
			drivers:    make([]domain.Driver, 0),
			bookings:   make([]domain.Booking, 0),
			driverPool: make([]string, 0),
		}

		return storage, nil
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	schema := &struct {
		Drivers    []domain.Driver  `json:"drivers"`
		Bookings   []domain.Booking `json:"orders"`
		DriverPool []string         `json:"driverPool"`
	}{}
	if err := json.Unmarshal(data, schema); err != nil {
		return nil, err
	}

	storage := &Storage{
		drivers:    schema.Drivers,
		bookings:   schema.Bookings,
		driverPool: schema.DriverPool,
	}
	return storage, nil
}

func (s *Storage) Save() error {
	schema := struct {
		Drivers    []domain.Driver  `json:"drivers"`
		Bookings   []domain.Booking `json:"orders"`
		DriverPool []string         `json:"driverPool"`
	}{
		Drivers:    s.drivers,
		Bookings:   s.bookings,
		DriverPool: s.driverPool,
	}

	data, err := json.Marshal(schema)
	if err != nil {
		return err
	}

	os.WriteFile(FILE_PATH, data, 0644)

	return nil
}
