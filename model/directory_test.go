package model

import(
    "testing"
    "path/filepath"
    "github.com/stevan/file-editor-server/util"
)

func TestDirectoryName(t *testing.T) {
    path, remover := util.CreateTempDir("foo")
    defer remover()

    directory, err := model.NewDirectoryModel(path)
    if err != nil {
        t.Logf("got an error: %s", err)
        t.FailNow()
    }

    if directory.Name() != path {
        t.Errorf(util.FormatTestMessage("got incorrect name"), directory.Name(), path)
    }
}

func TestEmptyDirectoryListing(t *testing.T) {
    path, remover := util.CreateTempDir("foo")
    defer remover()

    directory, err := model.NewDirectoryModel(path)
    if err != nil {
        t.Logf("got an error: %s", err)
        t.FailNow()
    }

    listing := directory.List()

    if listing.Size() != 0 {
        t.Errorf(util.FormatTestMessage("listing is not empty"), len(listing), 0)
    }

}

func TestDirectoryListing(t *testing.T) {
    path, paths, remover := util.CreatePopulatedTempDir("foo", []string{"bar.txt", "baz.txt"})
    defer remover()

    directory, err := model.NewDirectoryModel(path)
    if err != nil {
        t.Logf("got an error: %s", err)
        t.FailNow()
    }

    listing := directory.List()

    if listing.Size() != len(paths) {
        t.Errorf(util.FormatTestMessage("listing is not expected size"), len(listing), len(paths))
    }

    for i, p := range paths {
        name := listing.Get(i)["name"]
        path := filepath.Base(p)
        if name != path {
            t.Errorf(util.FormatTestMessage("wrong listing name"), name, path)
        }
    }
}



