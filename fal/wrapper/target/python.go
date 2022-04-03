package target

import (
	_ "embed"
	"fal/shared/fs"
)

var (
	//go:embed python/main.py
	_python_main []byte

	//go:embed python/pyproject.toml
	_python_proj []byte

	//go:embed python/poetry.lock
	_python_lock []byte
)

func WrapperPythonBuilder(wd *fs.Location) error {
	ld := wd.InnerLevel("python")

	err := ld.CreateDir()
	if err != nil {
		return err
	}

	f := fs.FileList{
		"main.py":        _python_main,
		"pyproject.toml": _python_proj,
		"poetry.lock":    _python_lock,
	}

	_, err = ld.CreateManyFiles(f)
	if err != nil {
		return err
	}

	return nil
}
