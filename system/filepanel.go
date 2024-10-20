package system

import (
	"encoding/json"
	"fmt"
	"os"
)

type FilePanelItem struct {
	Name       string `json:"name"`
	Size       string `json:"size"`
	IsDir      bool   `json:"isDir"`
	IsSelected bool   `json:"isSelected"`
}

type FilePanelContent struct {
	Items            []*FilePanelItem `json:"items"`
	CurrentItemIndex int              `json:"currentItemIndex"`
	PanelIndex       int              `json:"panelIndex"`
	CurrentPath      string           `json:"currentPath"`
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

func (f *FilePanel) MainAction() {
	item := f.content.Items[f.content.CurrentItemIndex]
	if item.IsDir {
		f.currentDirectory.Add(item.Name)
		f.UpdateContent()
	}
	f.content.CurrentItemIndex = 0
}

func (f *FilePanel) GoBack() {
	if len(f.currentDirectory.Items) == 0 {
		return
	}
	f.currentDirectory.Items = f.currentDirectory.Items[:len(f.currentDirectory.Items)-1]
	f.UpdateContent()
	f.content.CurrentItemIndex = 0
}

func (f *FilePanel) SetCurrentDirectory(path *Path) {
	f.currentDirectory = path
}

func (c *FilePanel) GetContentAsJson() string {
	bs, _ := json.Marshal(c.content)
	return string(bs)
}

func (c *FilePanel) UpdateContent() {

	c.content.CurrentPath = c.currentDirectory.String()

	files, err := os.ReadDir(c.currentDirectory.String())
	if err != nil {
		fmt.Println("Error reading directory", c.currentDirectory.String(), err)
		return
	}

	c.content.Items = make([]*FilePanelItem, 0)
	for _, file := range files {
		var item FilePanelItem
		item.Name = file.Name()
		if file.IsDir() {
			item.Size = "[DIR]"
		} else {
			fi, err := file.Info()
			if err == nil {
				item.Size = fmt.Sprintf("%d", fi.Size())
			} else {
				item.Size = "err"
			}

		}
		item.IsDir = file.IsDir()
		c.content.Items = append(c.content.Items, &item)
	}
}
