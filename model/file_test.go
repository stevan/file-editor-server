package model

import (
	"github.com/stevan/file-editor-server/util"
	"path/filepath"
	"testing"
)

func TestFileCreate(t *testing.T) {
	path, remover := util.CreateTempDir("foo")
	defer remover()

	file_path := filepath.Join(path, "foo.txt")

	file, err := CreateFileModel(file_path, "testing")
	if err != nil {
		t.Logf("got an error: %s", err)
		t.FailNow()
	}

	if file.Name() != file_path {
		t.Errorf(util.FormatTestMessage("got incorrect name"), file.Name(), file_path)
	}

	contents := file.Read()
	if contents != "testing" {
		t.Errorf(util.FormatTestMessage("file contents are incorrect"), contents, "testing")
	}

	file.Remove()
	if util.FileOrDirectoryExists(file.Name()) {
		t.Errorf("file should be gone, but it is not: %s", file.Name())
	}
}

func TestFileName(t *testing.T) {
	path, remover := util.CreateFileInTempDir("foo.txt")
	defer remover()

	file, err := NewFileModel(path)
	if err != nil {
		t.Logf("got an error: %s", err)
		t.FailNow()
	}

	if file.Name() != path {
		t.Errorf(util.FormatTestMessage("got incorrect name"), file.Name(), path)
	}
}

func TestFileReadWrite(t *testing.T) {
	path, remover := util.CreateFileInTempDir("foo.txt")
	defer remover()

	file, err := NewFileModel(path)
	if err != nil {
		t.Logf("got an error: %s", err)
		t.FailNow()
	}

	contents := file.Read()
	if contents != "" {
		t.Errorf(util.FormatTestMessage("file should be empty"), contents, path)
	}

	file.Write("testing")

	contents = file.Read()
	if contents != "testing" {
		t.Errorf(util.FormatTestMessage("file shouldn't be empty"), contents, path)
	}
}

func TestFileRemove(t *testing.T) {
	path, remover := util.CreateFileInTempDir("foo.txt")
	defer remover()

	file, err := NewFileModel(path)
	if err != nil {
		t.Logf("got an error: %s", err)
		t.FailNow()
	}

	file.Remove()
	if util.FileOrDirectoryExists(file.Name()) {
		t.Errorf("file should be gone, but it is not: %s", file.Name())
	}
}
