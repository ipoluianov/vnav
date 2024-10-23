package common

type Item struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	DisplayExt  string `json:"displayExt"`
	Size        string `json:"size"`
	DateTime    string `json:"datetime"`
	IsDir       bool   `json:"isDir"`
	IsSelected  bool   `json:"isSelected"`
	FullPath    string `json:"fullPath"`
	LinkPath    string `json:"linkPath"`
}
