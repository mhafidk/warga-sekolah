package frontend

import (
	"embed"
	"io/fs"
)

//go:embed all:build
var buildFS embed.FS

// GetFS returns the fs.FS representing the build directory
func GetFS() fs.FS {
	f, err := fs.Sub(buildFS, "build")
	if err != nil {
		panic(err)
	}
	return f
}
