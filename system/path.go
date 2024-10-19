package system

import "runtime"

type Path struct {
	TP    string   `json:"type"` // type
	Info  string   `json:"info"` // info
	Items []string `json:"items"`
}

// [type] [info] [...]
// type = localfs
// info = localfs = ["x", "w"]

func NewPath() *Path {
	var c Path
	c.TP = "localfs"
	if runtime.GOOS == "windows" {
		c.Info = "w"
	} else {
		c.Info = "x"
	}
	return &c
}

func (c *Path) Add(item string) {
	c.Items = append(c.Items, item)
}

func (c Path) String() string {
	result := ""
	for _, item := range c.Items {
		result += "/" + item
	}
	if result == "" {
		result = "/"

	}
	return result
}

func (c *Path) Clone() *Path {
	var result Path
	result.TP = c.TP
	result.Info = c.Info
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
