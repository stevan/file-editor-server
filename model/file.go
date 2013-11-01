package model

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func NewFileModel(path string) (*File, error) {

	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if info.IsDir() {
		return nil, errors.New(fmt.Sprintf("%s is a directory", path))
	}

	return &File{path}, nil
}

func CreateFileModel(path string, contents string) (*File, error) {

	dir := filepath.Dir(path)
	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			mkdir_err := os.MkdirAll(dir, 0755)
			if mkdir_err != nil {
				return nil, mkdir_err
			}
		} else {
			return nil, err
		}
	}

	write_err := ioutil.WriteFile(path, []byte(contents), 0600)
	if write_err != nil {
		return nil, write_err
	}

	return &File{path}, nil
}

type File struct {
	path string
}

func (handle *File) Name() string {
	return handle.path
}

func (handle *File) Read() string {
	contents, err := ioutil.ReadFile(handle.Name())
	if err != nil {
		panic(err)
	}
	return string(contents)
}

func (handle *File) Write(contents string) {
	err := ioutil.WriteFile(handle.Name(), []byte(contents), 0600)
	if err != nil {
		panic(err)
	}
}

func (handle *File) Remove() {
	err := os.Remove(handle.Name())
	if err != nil {
		panic(err)
	}
}
