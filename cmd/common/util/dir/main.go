package dir

import (
	"os"
	"tomator/cmd/common/helper"
)

func Init(dir string) *Model {

	m := &Model{
		Path: dir,
	}
	m.Files = m.getFiles()
	return m

}

func (_m *Model) getFiles() (m map[string]string) {
	m = map[string]string{}

	files, err := os.ReadDir(_m.Path)
	helper.CheckError(err, true)
	for _, file := range files {
		if file.IsDir() {
			continue
		}

	}

	return
}
