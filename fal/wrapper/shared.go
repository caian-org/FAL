package wrapper

import (
	_ "embed"
	"fal/shared/fs"
)

var (
	//go:embed lib/main._go
	_libfal_main []byte

	//go:embed lib/go._mod
	_libfal_mod []byte

	//go:embed lib/go._sum
	_libfal_sum []byte

	//go:embed lib/_Makefile
	_libfal_make []byte
)

func InitSharedLib(buildDir *fs.Location) error {
	wd := buildDir.InnerLevel("lib")

	err := wd.CreateDir()
	if err != nil {
		return err
	}

	f := fs.FileList{
		"main.go":  _libfal_main,
		"go.mod":   _libfal_mod,
		"go.sum":   _libfal_sum,
		"Makefile": _libfal_make,
	}

	_, err = wd.CreateManyFiles(f)
	if err != nil {
		return err
	}

	return nil
}
