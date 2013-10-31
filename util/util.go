package util

import(
    "os"
    "mime"
    "strings"
    "errors"
    "path/filepath"

    "github.com/stevan/httpapp"
)

// ------------------------------------------------------
// HTTP Utils
// ------------------------------------------------------

func SlurpRequestBody (e *httpapp.Env) string {
    content := make([]byte, e.Request.ContentLength)
    _, err := e.Request.Body.Read(content)
    if err != nil { panic(err) }
    return string(content)
}

// Header utils

func MediaType (name string, params ... string) string {
    if len(params) % 2 != 0 {
        panic(errors.New("params must be an even sized list"))
    }
    p := make(map[string]string, len(params) / 2)
    for i := 0; i < len(params); i += 2 {
        p[ params[i] ] = params[ i + 1 ]
    }
    return mime.FormatMediaType( name, p )
}

func GuessMediaType (name string) string {
    parts := strings.Split(name, ".")
    return mime.TypeByExtension( "." + parts[ len(parts) - 1 ] )
}

// ------------------------------------------------------
// Test Utils
// ------------------------------------------------------

func FormatTestMessage (msg string) string {
    return msg + "\n     got: %v\nexpected: %v"
}

// ------------------------------------------------------
// TempDir/File Utils
// ------------------------------------------------------

/*
 * TODO:
 * - replace this with TempDir and TempFile from io/ioutil
 */

func CreateTempDir (name string) (string, func()) {
    path := filepath.Join(os.TempDir(), name)
    os.Mkdir(path, 0755)
    return path, func () { os.Remove(path) }
}

func CreateFileInTempDir (name string) (string, func()) {
    path := filepath.Join(os.TempDir(), name)
    _, err := os.Create(path)
    if err != nil { panic(err) }
    return path, func () { os.Remove(path) }
}

func CreatePopulatedTempDir (name string, files []string) (string, []string, func()) {
    path := filepath.Join(os.TempDir(), name)
    os.Mkdir(path, 0755)
    paths := make([]string, len(files))
    for i, f := range files {
        paths[i] = filepath.Join(path, f)
        _, err := os.Create(paths[i])
        if err != nil { panic(err) }
    }
    return path, paths, func () { os.RemoveAll(path) }
}

func FileOrDirectoryExists (name string) bool {
    if _, err := os.Stat(name); err != nil {
        if os.IsNotExist(err) {
            return false
        }
    }
    return true
}


