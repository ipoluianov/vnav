package system

import (
	"runtime"
	"sync"
	"vnav/drivers/driver_localfs"
)

type System struct {
	drivers    *Drivers
	filePanels [2]*FilePanel
	mtx        sync.Mutex
}

func NewSystem() *System {
	var c System
	c.drivers = NewDrivers()

	if runtime.GOOS == "windows" {
		drives, err := driver_localfs.GetDrives()
		if err == nil {
			for _, drive := range drives {
				c.drivers.Register(driver_localfs.NewDriverLocalFS(drive))
			}
		}

		d0, _ := c.drivers.GetDriver("C:")
		d1, _ := c.drivers.GetDriver("D:")

		c.filePanels[0] = NewFilePanel(0, d0)
		c.filePanels[1] = NewFilePanel(1, d1)
	}

	if runtime.GOOS == "linux" {
		c.drivers.Register(driver_localfs.NewDriverLocalFS(""))
		d, _ := c.drivers.GetDriver("")

		c.filePanels[0] = NewFilePanel(0, d)
		c.filePanels[1] = NewFilePanel(1, d)
	}
	if runtime.GOOS == "darwin" {
		c.drivers.Register(driver_localfs.NewDriverLocalFS(""))
		d, _ := c.drivers.GetDriver("")

		c.filePanels[0] = NewFilePanel(0, d)
		c.filePanels[1] = NewFilePanel(1, d)
	}

	return &c
}

func (c *System) GetFilePanelContentAsJson(index int) string {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	return c.filePanels[index].GetContentAsJson()
}

func (c *System) GetFilePanelContent(index int) *FilePanelContent {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	return &c.filePanels[index].content
}

func (c *System) SetCurrentItemIndex(panelIndex int, index int) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.filePanels[panelIndex].SetCurrentItemIndex(index)
}

func (c *System) UpdateContent(panelIndex int) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.filePanels[panelIndex].UpdateContent()
}

func (c *System) MainAction(panelIndex int) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.filePanels[panelIndex].MainAction()
}

func (c *System) GoBack(panelIndex int) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.filePanels[panelIndex].GoBack()
}

func (c *System) CreateDirectory(name string, panelIndex int) error {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	return c.filePanels[panelIndex].CreateDirectory(name)
}
