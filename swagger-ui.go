package go_swagger_ui_resource

import (
	"embed"
	"io/fs"
	"path"
)

//go:embed static
var data embed.FS

type _FS struct{}

var FS _FS

func (_FS) Open(name string) (fs.File, error) {
	return data.Open(path.Join("static", name))
}
