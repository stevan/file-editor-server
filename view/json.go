package view

import (
    "encoding/json"
)

func NewJSONView () *JSONView {
    return new(JSONView)
}

type JSONView struct {}

func (v *JSONView) Render (data interface {}) []byte {
    json, err := json.MarshalIndent(data, "", "  ")
    if err != nil { panic(err) }
    return json
}

