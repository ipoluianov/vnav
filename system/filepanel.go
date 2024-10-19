package system

import (
	"encoding/json"
	"fmt"
	"os"
)

type FilePanelItem struct {
	Name       string `json:"name"`
	Size       int64  `json:"size"`
	IsDir      bool   `json:"isDir"`
	IsSelected bool   `json:"isSelected"`
}

type FilePanelContent struct {
	Items            []*FilePanelItem `json:"items"`
	CurrentItemIndex int              `json:"currentItemIndex"`
	PanelIndex       int              `json:"panelIndex"`
}

type FilePanel struct {
	currentDirectory *Path
	content          FilePanelContent
}

func NewFilePanel(index int) *FilePanel {
	var c FilePanel
	c.content.PanelIndex = index
	c.currentDirectory = NewPath()
	c.content.Items = make([]*FilePanelItem, 0)
	c.content.CurrentItemIndex = 0
	c.content.PanelIndex = index
	c.UpdateContent()
	return &c
}

func (f *FilePanel) SetCurrentItemIndex(index int) {
	if index < 0 {
		index = 0
	}
	if index >= len(f.content.Items) {
		index = len(f.content.Items) - 1
	}
	f.content.CurrentItemIndex = index
}

func (f *FilePanel) SetCurrentDirectory(path *Path) {
	f.currentDirectory = path
}

func (c *FilePanel) GetContentAsJson() string {
	bs, _ := json.Marshal(c.content)
	return string(bs)
}

func (c *FilePanel) UpdateContent() {
	files, err := os.ReadDir(c.currentDirectory.String())
	if err != nil {
		fmt.Println("Error reading directory", c.currentDirectory.String(), err)
		return
	}

	c.content.Items = make([]*FilePanelItem, 0)
	for _, file := range files {
		c.content.Items = append(c.content.Items, &FilePanelItem{Name: file.Name(), IsDir: file.IsDir()})
	}
}
