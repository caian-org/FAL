package wrapper

import (
	_ "embed"
	"fal/util"
)

var (
	//go:embed python/main.py
	_python_main []byte

	//go:embed python/pyproject.toml
	_python_pyproj []byte

	//go:embed python/poetry.lock
	_python_lock []byte
)

func BuildPythonWrapper(wd *util.Location) error {
	ld := wd.InnerLevel("python")

	err := ld.CreateDir()
	if err != nil {
		return err
	}

	_, err = ld.CreateFile("main.py", _python_main)
	if err != nil {
		return err
	}

	return nil
}
