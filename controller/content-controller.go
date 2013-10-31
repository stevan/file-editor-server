package controller

import (
    "fmt"
    "os"
    //"log"
    "path/filepath"
    "net/http"

    "github.com/stevan/file-editor-server/model"
    "github.com/stevan/file-editor-server/view"
    "github.com/stevan/file-editor-server/util"

    "github.com/stevan/httpapp"
)

func NewContentController (base string) *ContentController {
    return &ContentController{base}
}

type ContentController struct {
    base string
}

/*
 * Directories:
 *   GET    list contents
 * Files:
 *   POST   create file (and containing directory if needed)
 *   PUT    update file
 *   GET    fetch file
 *   DELETE remove file
 */

func (c *ContentController) ServeHTTP (w http.ResponseWriter, r *http.Request) {
    c.Call(httpapp.NewEnv(r)).WriteTo(w)
}

func (c *ContentController) Call (e *httpapp.Env) *httpapp.Response {
    path      := c.getFullPath(e.Request.URL.Path)
    info, err := os.Stat(path)
    if err != nil {
        if os.IsNotExist(err) && e.Request.Method == "POST" {
            if c.isDir(path) {
                return c.handleDirectoryPOST(path, e)
            } else {
                return c.handleFilePOST(path, e)
            }
        } else {
            panic(err)
        }
    } else {
        if info.IsDir() {
            return c.handleDirectory(path, e)
        } else {
            return c.handleFile(path, e)
        }
    }
}

// operations on paths

func (c *ContentController) handleDirectory (path string, e *httpapp.Env) *httpapp.Response {
    d, err := model.NewDirectoryModel(path)
    if err != nil { panic(err) }

    if e.Request.Method == "GET" {
        resp := httpapp.NewResponse(http.StatusOK)
        resp.Headers.Add("Content-Type", util.MediaType( "application/json", "charset", "utf-8" ))
        resp.Body.Write(view.NewJSONView().Render(d.List()))
        return resp
    } else {
        resp := httpapp.NewResponse(http.StatusMethodNotAllowed)
        resp.Body.WriteString(fmt.Sprintf("The %s method is not supported on directories", e.Request.Method))
        return resp
    }
}

func (c *ContentController) handleFile (path string, e *httpapp.Env) *httpapp.Response {
    f, err := model.NewFileModel(path)
    if err != nil { panic(err) }

    if e.Request.Method == "GET" {
        return c.handleFileGET(f, e)
    } else if e.Request.Method == "PUT" {
        return c.handleFilePUT(f, e)
    } else if e.Request.Method == "DELETE" {
        return c.handleFileDELETE(f, e)
    } else {
        resp := httpapp.NewResponse(http.StatusMethodNotAllowed)
        resp.Body.WriteString(fmt.Sprintf("The %s method is not supported on existing files", e.Request.Method))
        return resp
    }
}

func (c *ContentController) handleFilePOST (path string, e *httpapp.Env) *httpapp.Response {
    f, err := model.CreateFileModel(path, util.SlurpRequestBody(e))
    if err != nil { panic(err) }
    resp := httpapp.NewResponse(http.StatusNoContent)
    resp.Headers.Add("Location",     e.Request.URL.Path)
    resp.Headers.Add("Content-Type", util.GuessMediaType( f.Name() ))
    return resp
}

func (c *ContentController) handleDirectoryPOST (path string, e *httpapp.Env) *httpapp.Response {
    resp := httpapp.NewResponse(http.StatusMethodNotAllowed)
    resp.Body.WriteString("The POST method is not supported on directories")
    return resp
}

// operate on file instances

func (c *ContentController) handleFileGET (f *model.File, e *httpapp.Env) *httpapp.Response {
    resp := httpapp.NewResponse(http.StatusOK)
    resp.Headers.Add("Content-Type", util.GuessMediaType( f.Name() ))
    resp.Body.WriteString(f.Read())
    return resp
}

func (c *ContentController) handleFilePUT (f *model.File, e *httpapp.Env) *httpapp.Response {
    f.Write(util.SlurpRequestBody(e))
    return httpapp.NewResponse(http.StatusNoContent)
}

func (c *ContentController) handleFileDELETE (f *model.File, e *httpapp.Env) *httpapp.Response {
    f.Remove()
    return httpapp.NewResponse(http.StatusNoContent)
}


// utils ...

func (c *ContentController) isDir (path string) bool {
    return filepath.Dir(path) == path
}

func (c *ContentController) getFullPath (path string) string {
    return filepath.Join(c.base, filepath.Clean(path))
}



