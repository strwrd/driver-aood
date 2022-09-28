package business

import "github.com/strwrd/driver-aood/src/domain"

func (c *Case) RegisterDriver(id string) (*domain.Driver, error) {
	// create driver, returning error when driver id is duplicated
	d := &domain.Driver{ID: id}
	if err := c.Storage.InsertDriver(d); err != nil {
		return nil, err
	}

	// insert driver id into available driver pool
	if err := c.Storage.InsertDriverPool(id); err != nil {
		return nil, err
	}
	return d, nil
}
