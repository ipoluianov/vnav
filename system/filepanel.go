package system

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"vnav/common"
)

type FilePanelContent struct {
	Items            []*common.Item `json:"items"`
	CurrentItemIndex int            `json:"currentItemIndex"`
	PanelIndex       int            `json:"panelIndex"`
	CurrentPath      string         `json:"currentPath"`
}

type FilePanelContext struct {
	CurrentItemIndex int
}

type FilePanel struct {
	driver           common.Driver
	currentDirectory *common.Path
	content          FilePanelContent
	contextStack     []FilePanelContext
}

func NewFilePanel(index int, driver common.Driver) *FilePanel {
	var c FilePanel
	c.driver = driver
	c.content.PanelIndex = index
	c.currentDirectory = driver.GetRoot()
	c.content.Items = make([]*common.Item, 0)
	c.content.CurrentItemIndex = 0
	c.content.PanelIndex = index
	c.contextStack = append(c.contextStack, FilePanelContext{CurrentItemIndex: 0})
	c.UpdateContent()
	return &c
}

func (c *FilePanel) SetCurrentItemIndex(index int) {
	if index < 0 {
		index = 0
	}
	if index >= len(c.content.Items) {
		index = len(c.content.Items) - 1
	}
	c.content.CurrentItemIndex = index
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
		f.openFile(item.FullPath)
	}
}

func (f *FilePanel) openFile(filePath string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", filePath)
	case "darwin":
		cmd = exec.Command("open", filePath)
	case "linux":
		cmd = exec.Command("xdg-open", filePath)
	default:
		return fmt.Errorf("unsupported platform")
	}
	return cmd.Start()
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

func (f *FilePanel) SetCurrentDirectory(path *common.Path) {
	f.currentDirectory = path
}

func (c *FilePanel) GetContentAsJson() string {
	bs, _ := json.Marshal(c.content)
	return string(bs)
}

func (c *FilePanel) UpdateContent() error {
	if true {
		items, err := c.driver.ReadDir(c.currentDirectory)
		if err != nil {
			return err
		}

		c.content.CurrentPath = c.driver.PathToString(c.currentDirectory)
		c.content.Items = make([]*common.Item, 0)

		if len(c.currentDirectory.Items) > 0 {
			c.content.Items = append(c.content.Items, &common.Item{Name: "..", DisplayName: "[..]", Size: "[DIR]", IsDir: true})
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
