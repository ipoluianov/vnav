package system

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

type FilePanelContext struct {
	CurrentItemIndex int
}

type FilePanel struct {
	currentDirectory *Path
	content          FilePanelContent
	contextStack     []FilePanelContext
}

func NewFilePanel(index int) *FilePanel {
	var c FilePanel
	c.content.PanelIndex = index
	c.currentDirectory = NewPath()
	c.content.Items = make([]*FilePanelItem, 0)
	c.content.CurrentItemIndex = 0
	c.content.PanelIndex = index
	c.contextStack = append(c.contextStack, FilePanelContext{CurrentItemIndex: 0})
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
		if item.Name == ".." {
			f.GoBack()
		} else {
			f.GoToDirectory()
		}
	} else {
		fmt.Println("Selected file", item.Name)
	}
}

func (f *FilePanel) GoToDirectory() {
	item := f.content.Items[f.content.CurrentItemIndex]

	f.currentDirectory.Items = append(f.currentDirectory.Items, item.Name)

	err := f.UpdateContent()
	if err != nil {
		f.currentDirectory.Items = f.currentDirectory.Items[:len(f.currentDirectory.Items)-1]
		f.UpdateContent()
		return
	}

	if len(f.contextStack) > 0 {
		f.contextStack[len(f.contextStack)-1].CurrentItemIndex = f.content.CurrentItemIndex
	}

	f.content.CurrentItemIndex = 0
	f.contextStack = append(f.contextStack, FilePanelContext{CurrentItemIndex: 0})
}

func (f *FilePanel) GoBack() {
	if len(f.currentDirectory.Items) == 0 {
		return
	}
	lastPath := f.currentDirectory.Clone()
	f.currentDirectory.Items = f.currentDirectory.Items[:len(f.currentDirectory.Items)-1]
	err := f.UpdateContent()
	if err != nil {
		f.currentDirectory = lastPath
		f.UpdateContent()
		return
	}

	if len(f.contextStack) > 1 {
		f.contextStack = f.contextStack[:len(f.contextStack)-1]
		context := f.contextStack[len(f.contextStack)-1]
		f.content.CurrentItemIndex = context.CurrentItemIndex
	}
}

func (f *FilePanel) SetCurrentDirectory(path *Path) {
	f.currentDirectory = path
}

func (c *FilePanel) GetContentAsJson() string {
	bs, _ := json.Marshal(c.content)
	return string(bs)
}

func (c *FilePanel) UpdateContent() error {
	c.content.CurrentPath = c.currentDirectory.String()

	fmt.Println("Reading directory", c.currentDirectory.String())

	files, err := os.ReadDir(c.currentDirectory.String())
	if err != nil {
		fmt.Println("Error reading directory", c.currentDirectory.String(), err)
		return err
	}

	c.content.Items = make([]*FilePanelItem, 0)
	items := make([]*FilePanelItem, 0)
	for _, file := range files {
		var item FilePanelItem
		item.Name = file.Name()

		infoAvailable := true

		fi, err := file.Info()
		if err != nil {
			infoAvailable = false
		}

		path := c.currentDirectory.String() + "/" + file.Name()
		lsInfo, err := os.Lstat(path)
		if err != nil {
			infoAvailable = false
		}

		if infoAvailable {

			if fi.IsDir() {
				item.IsDir = true
			} else {
				if lsInfo.Mode()&os.ModeSymlink != 0 {
					linkPath, err := filepath.EvalSymlinks(path)
					if err != nil {
						infoAvailable = false
					} else {
						targetInfo, err := os.Stat(linkPath)
						if err != nil {
							infoAvailable = false
						} else {
							if targetInfo.IsDir() {
								item.IsDir = true
							}
						}
					}
				} else {
					item.IsDir = false
					item.Size = formatFileSize(fi.Size())
				}
			}

			if item.IsDir {
				item.Size = "[DIR]"
			}
		}

		items = append(items, &item)
	}

	if len(c.currentDirectory.Items) > 0 {
		c.content.Items = append(c.content.Items, &FilePanelItem{Name: "..", Size: "[DIR]", IsDir: true})
	}

	for _, item := range items {
		if item.IsDir {
			c.content.Items = append(c.content.Items, item)
		}
	}

	for _, item := range items {
		if !item.IsDir {
			c.content.Items = append(c.content.Items, item)
		}
	}

	return nil
}

func formatFileSize(bytes int64) string {
	result := formatNumber(fmt.Sprint(bytes))
	return result
}

func formatNumber(s string) string {
	var result []string
	n := len(s)
	for i := n; i > 0; i -= 3 {
		start := i - 3
		if start < 0 {
			start = 0
		}
		result = append([]string{s[start:i]}, result...)
	}
	return strings.Join(result, " ")
}
