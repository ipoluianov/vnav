package system

import (
	"errors"
	"sync"
	"vnav/common"
)

type Drivers struct {
	mtx     sync.Mutex
	drivers map[string]common.Driver
}

func NewDrivers() *Drivers {
	var c Drivers
	c.drivers = make(map[string]common.Driver)
	return &c
}

func (c *Drivers) Register(driver common.Driver) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.drivers[driver.Id()] = driver
}

func (c *Drivers) GetDriver(id string) (common.Driver, error) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	if _, ok := c.drivers[id]; !ok {
		return nil, errors.New("driver not found")
	}
	return c.drivers[id], nil
}
