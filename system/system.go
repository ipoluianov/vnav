package system

import (
	"sync"
)

type System struct {
	filePanels [2]*FilePanel
	mtx        sync.Mutex
}

func NewSystem() *System {
	var c System
	c.filePanels[0] = NewFilePanel(0)
	c.filePanels[1] = NewFilePanel(1)
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
