package file

import (
	"os"
)

type File struct {
	Name          string
	BaseName      string
	IsValid       bool
	Path          string
	Ext           string
	ExtType       *Ext
	Content       string
	FileWithNoExt string
	New           *os.File
}
type Ext struct {
	JSON string
	YAML string
	TEXT string
}
