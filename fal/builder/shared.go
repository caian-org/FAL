package builder

import (
	_ "embed"
	"fal/util"
)

var (
	//go:embed shared/fal.go
	_libfal_main []byte

	//go:embed shared/_go.mod
	_libfal_mod []byte

	//go:embed shared/_go.sum
	_libfal_sum []byte

	//go:embed shared/Makefile
	_libfal_make []byte
)

func InitSharedLib(buildDir *util.Location) error {
	wd := buildDir.InnerLevel("shared")

	err := wd.CreateDir()
	if err != nil {
		return err
	}

	f := util.FileList{
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
