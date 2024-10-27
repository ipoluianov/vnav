package common

type Driver interface {
	Id() string
	ReadDir(path *Path) ([]*Item, error)
	PathToString(path *Path) string
	GetRoot() *Path
	Separator() string
	CreateDirectory(path *Path, name string) error
	Remove(path *Path, name string) error
}
