package controller

import(
    "testing"

    "path/filepath"
    "encoding/json"
    "net/http"
    "net/http/httptest"

    "github.com/stevan/file-editor-server/model"
    "github.com/stevan/file-editor-server/util"
)

func TestDirectoryListing (t *testing.T) {
    path, paths, remover := util.CreatePopulatedTempDir("foo", []string{"bar.txt", "baz.txt"})
    defer remover()

    contentserver := NewContentController( path )

    req, err := http.NewRequest("GET", "http://example.com/", nil)
    if err != nil {
        t.Logf("got an error: %s", err)
        t.FailNow()
    }

    w := httptest.NewRecorder()
    contentserver.ServeHTTP(w, req)

    if w.Code != 200 {
        t.Errorf(util.FormatTestMessage("got wrong status"), w.Code, 200)
    }

    listing := make(model.DirectoryListing, len(paths))
    body    := w.Body.String()

    err = json.Unmarshal([]byte(body), &listing)
    if err != nil {
        t.Errorf("got an error: %s", err)
    }

    if listing.Size() != len(paths) {
        t.Errorf(util.FormatTestMessage("got wrong size of listing"), listing.Size(), len(paths))
    }

    for i, p := range paths {
        name := listing.Get(i)["name"]
        path := filepath.Base(p)
        if name != path {
            t.Errorf(util.FormatTestMessage("wrong listing name"), name, path)
        }
    }
}
