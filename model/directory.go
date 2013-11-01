package model

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func NewDirectoryModel(path string) (*Directory, error) {

	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return nil, errors.New(fmt.Sprintf("%s is not a directory", path))
	}

	return &Directory{path}, nil
}

type Directory struct {
	path string
}

func (handle *Directory) Name() string {
	return handle.path
}

func (handle *Directory) List() DirectoryListing {
	contents, err := ioutil.ReadDir(handle.path)
	if err != nil {
		panic(err)
	}
	return createDirectoryListing(contents)
}

// private constructor

func createDirectoryListing(contents []os.FileInfo) DirectoryListing {
	listing := make(DirectoryListing, len(contents))

	for idx, info := range contents {
		if info.IsDir() {
			listing[idx] = DirectoryItem{"name": info.Name(), "type": "dir"}
		} else {
			listing[idx] = DirectoryItem{"name": info.Name(), "type": "file"}
		}
	}

	return listing
}

type DirectoryItem map[string]string
type DirectoryListing []DirectoryItem

func (listing DirectoryListing) Size() int {
	return len(listing)
}

func (listing DirectoryListing) Get(i int) map[string]string {
	return listing[i]
}
