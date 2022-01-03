package wrapper

import (
	_ "embed"
	"fal/util"
)

var (
	//go:embed javascript/index.js
	_js_main []byte

	//go:embed javascript/package.json
	_js_package []byte

	//go:embed javascript/package-lock.json
	_js_lock []byte
)

func BuildJavaScriptWrapper(wd *util.Location) error {
	ld := wd.InnerLevel("javascript")

	err := ld.CreateDir()
	if err != nil {
		return err
	}

	f := util.FileList{
		"package.json":      _js_package,
		"package-lock.json": _js_lock,
	}

	_, err = ld.CreateManyFiles(f)
	if err != nil {
		return err
	}

	return nil
}
