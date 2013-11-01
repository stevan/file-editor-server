package model

import (
	"errors"
	"os/exec"
	"path/filepath"
)

func NewVersionedFileModel(path string) (*VersionedFile, error) {
	f, err := NewFileModel(path)
	if err != nil {
		return nil, err
	}
	return &VersionedFile{file: f, is_new: false}, nil
}

func CreateVersionedFileModel(path string, contents string) (*VersionedFile, error) {
	f, err := CreateFileModel(path, contents)
	if err != nil {
		return nil, err
	}
	return &VersionedFile{file: f, is_new: true}, nil
}

type VersionedFile struct {
	file   *File
	is_new bool
}

func (handle *VersionedFile) Name() string {
	return handle.file.path
}

func (handle *VersionedFile) Read() string {
	return handle.file.Read()
}

func (handle *VersionedFile) Write(contents string) {
	handle.file.Write(contents)
}

func (handle *VersionedFile) Remove() {
	handle.file.Remove()
}

func (handle *VersionedFile) CommitFor(author string) string {
	all_out := ""

	working_dir := filepath.Dir(handle.Name())

	if handle.is_new == true {
		cmd := exec.Command("git", "add", handle.Name())
		cmd.Dir = working_dir
		out, err := cmd.CombinedOutput()
		all_out = all_out + string(out)
		if err != nil {
			panic(errors.New("Error: <" + err.Error() + "> Message: " + string(out)))
		}
	}

	cmd := exec.Command("git", "commit", "--author", author, "-am", "commiting "+handle.Name())
	cmd.Dir = working_dir
	out, err := cmd.CombinedOutput()
	all_out = all_out + string(out)
	if err != nil {
		panic(errors.New("Error: <" + err.Error() + "> Message: " + string(out)))
	}

	return all_out
}
