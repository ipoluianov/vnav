package common

type Driver interface {
	Id() string
	ReadDir(path *Path) ([]*Item, error)
	PathToString(path *Path) string
	GetRoot() *Path
	Separator() string
}
