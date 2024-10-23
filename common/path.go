package common

import "runtime"

type Path struct {
	Driver   string   `json:"type"`     // type
	RootName string   `json:"rootName"` // rootName
	Items    []string `json:"items"`    // items
}

// [type] [info] [...]
// type = localfs
// info = localfs = ["x", "w"]

func NewPath() *Path {
	var c Path
	c.Driver = "localfs"
	c.RootName = ""
	if runtime.GOOS == "windows" {
		c.RootName = "C:\\"
	}
	return &c
}

func (c *Path) Add(item string) {
	c.Items = append(c.Items, item)
}

func (c Path) String(separator string) string {
	result := ""
	for _, item := range c.Items {
		result += separator + item
	}
	if result == "" {
		result = separator
	}
	return result
}

func (c *Path) Clone() *Path {
	var result Path
	result.Driver = c.Driver
	result.Items = append(result.Items, c.Items...)
	return &result
}

func (c Path) Last() string {
	if len(c.Items) == 0 {
		return ""
	}
	return c.Items[len(c.Items)-1]
}

func (c Path) Parent() *Path {
	if len(c.Items) == 0 {

		return NewPath()
	}
	return &Path{Items: c.Items[:len(c.Items)-1]}
}

func (c Path) IsRoot() bool {
	return len(c.Items) == 0
}
