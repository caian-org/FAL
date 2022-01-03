package wrapper

import (
	_ "embed"
	"fal/util"
)

var (
	//go:embed python/fal.py
	_python_main []byte
)

func BuildPythonWrapper(wd *util.Location) error {
	ld := wd.InnerLevel("python")

	err := ld.CreateDir()
	if err != nil {
		return err
	}

	_, err = ld.CreateFile("fal.py", _python_main)
	if err != nil {
		return err
	}

	return nil
}
