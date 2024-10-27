package driver_localfs

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
	"vnav/common"
)

type DriverLocalFS struct {
	id string
}

func NewDriverLocalFS(id string) *DriverLocalFS {
	var c DriverLocalFS
	c.id = id
	return &c
}

func GetDrives() ([]string, error) {
	var drives []string
	return drives, nil
}

func (c *DriverLocalFS) Id() string {
	return c.id
}

func (c *DriverLocalFS) GetRoot() *common.Path {
	p := common.NewPath()
	p.RootName = c.id
	p.Items = make([]string, 0)
	p.Driver = c.id
	return p
}

func (c *DriverLocalFS) Separator() string {
	return "/"
}

func (c *DriverLocalFS) PathToString(path *common.Path) string {
	result := c.id + path.String(c.Separator())
	return result
}

func (c *DriverLocalFS) ReadDir(path *common.Path) ([]*common.Item, error) {
	var items []*common.Item

	dir := path.String(c.Separator())

	if path.RootName != "" {
		dir = path.RootName + dir
	}

	dateTime := time.Unix(0, 0)

	fmt.Println("Reading directory", dir)

	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory", dir, err)
		return nil, err
	}

	for _, file := range files {
		var item common.Item
		item.Name = file.Name()

		infoAvailable := true

		fi, err := file.Info()
		if err != nil {
			infoAvailable = false
		}

		path := dir + c.Separator() + file.Name()
		if dir == c.Separator() {
			path = dir + file.Name()
		}

		item.FullPath = path

		lsInfo, err := os.Lstat(path)
		if err != nil {
			infoAvailable = false
		}

		if infoAvailable {
			dateTime = fi.ModTime()

			if fi.IsDir() {
				item.IsDir = true
			} else {
				if lsInfo.Mode()&os.ModeSymlink != 0 {
					linkPath, err := filepath.EvalSymlinks(path)
					if err != nil {
						infoAvailable = false
					} else {
						item.LinkPath = linkPath
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
				item.Size = "<DIR>"
				item.DisplayName = "[" + item.Name + "]"
			} else {
				item.DisplayName = item.Name
				item.DisplayExt = filepath.Ext(item.Name)
				if len(item.DisplayExt) > 0 {
					item.DisplayName = item.DisplayName[:len(item.DisplayName)-len(item.DisplayExt)]
					item.DisplayExt = item.DisplayExt[1:]
				}
			}
		}

		item.DateTime = dateTime.Format("2006-01-02 15:04:05")

		items = append(items, &item)
	}

	return items, nil
}

func (c *DriverLocalFS) CreateDirectory(path *common.Path, name string) error {
	dir := path.String(c.Separator())

	if path.RootName != "" {
		dir = path.RootName + dir
	}

	err := os.Mkdir(dir+c.Separator()+name, 0755)
	if err != nil {
		return err
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
