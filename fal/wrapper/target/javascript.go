package target

import (
	_ "embed"
	"fal/shared/fs"
)

var (
	//go:embed javascript/index.js
	_javascript_main []byte

	//go:embed javascript/package.json
	_javascript_package []byte

	//go:embed javascript/package-lock.json
	_javascript_lock []byte
)

func WrapperJavascriptBuilder(wd *fs.Location) error {
	ld := wd.InnerLevel("javascript")

	err := ld.CreateDir()
	if err != nil {
		return err
	}

	f := fs.FileList{
		"index.js":          _javascript_main,
		"package.json":      _javascript_package,
		"package-lock.json": _javascript_lock,
	}

	_, err = ld.CreateManyFiles(f)
	if err != nil {
		return err
	}

	return nil
}
