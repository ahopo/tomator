package util

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"tomator/cmd/common/helper"

	"gopkg.in/yaml.v2"
)

func InitFile(file_name string) *File {
	f := &File{
		Name: file_name,
	}
	f.IsValid = f.valid()
	f.BaseName = strings.ReplaceAll(filepath.Base(file_name), filepath.Ext(file_name), "")
	if !f.valid() {
		log.Fatal("File does not exist.\nFILE:", file_name)
	}
	abs, _ := filepath.Abs(file_name)
	f.Path = strings.ReplaceAll(abs, filepath.Base(file_name), "")
	f.Ext = f.getExt()
	f.Content = f.read()
	f.FileWithNoExt = filepath.Join(f.Path, f.BaseName)
	f.ExtType = initExtType()

	return f
}
func initExtType() (ext *Ext) {
	ext = &Ext{
		TEXT: "txt",
		YAML: "yml",
		JSON: "json",
	}
	return ext
}

// private func
func (f *File) read() string {

	content, error := os.ReadFile(f.Name)
	if error != nil {

		log.Fatal(error)
	}
	return string(content)
}
func (f *File) getExt() string {

	ret := strings.ToLower(filepath.Ext(f.Name))
	if ret == ".yaml" {
		return ".yml"
	}
	return ret
}
func (f *File) valid() bool {
	if _, err := os.Stat(f.Name); err == nil {
		return true
	}
	return false
}

// TODO : Public func
func (_f *File) CreateUsingExistingName(ext string) string {
	theNewFile := strings.Join([]string{_f.FileWithNoExt, ext}, ".")
	f, err := os.Create(theNewFile)
	helper.CheckError(err, true)
	_f.New = f
	return theNewFile
}
func (_f *File) Write(content string) bool {
	f := _f.New
	_, err := f.WriteString(content)
	if err != nil {
		helper.CheckError(err, false)
		f.Close()
		return false
	}
	return true
}

func SturctToYML(_struct interface{}) string {
	data, err := yaml.Marshal(_struct)
	helper.CheckError(err, false)
	return string(data)
}
