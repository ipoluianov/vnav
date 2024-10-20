package system

import "fmt"

type System struct {
	filePanels [2]*FilePanel
}

func NewSystem() *System {
	var c System
	c.filePanels[0] = NewFilePanel(0)
	c.filePanels[1] = NewFilePanel(1)
	return &c
}

func (c *System) GetFilePanelContentAsJson(index int) string {
	return c.filePanels[index].GetContentAsJson()
}

func (c *System) GetFilePanelContent(index int) *FilePanelContent {
	return &c.filePanels[index].content
}

func (c *System) SetCurrentItemIndex(panelIndex int, index int) {
	c.filePanels[panelIndex].SetCurrentItemIndex(index)
}

func (c *System) UpdateContent(panelIndex int) {
	fmt.Println("UpdateContent", panelIndex)
	c.filePanels[panelIndex].UpdateContent()
}

func (c *System) MainAction(panelIndex int) {
	c.filePanels[panelIndex].MainAction()
}

func (c *System) GoBack(panelIndex int) {
	c.filePanels[panelIndex].GoBack()
}
