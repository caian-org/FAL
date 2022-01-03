package util

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type FileList map[string][]byte

type Location struct {
	Path string
}

func (c Location) join(d string) string {
	return filepath.Join(c.Path, d)
}

func NewLocation(path string) *Location {
	return &Location{path}
}

func (c Location) InnerLevel(d string) *Location {
	return &Location{Path: c.join(d)}
}

func (c Location) CreateDir() error {
	return os.MkdirAll(c.Path, 0755)
}

func (c Location) CreateFile(filename string, data []byte) (string, error) {
	filepath := c.join(filename)
	err := os.WriteFile(filepath, data, 0644)
	if err != nil {
		return "", err
	}

	return filepath, nil
}

func (c Location) CreateManyFiles(files FileList) ([]string, error) {
	var created []string
	var failed []string

	for file, data := range files {
		filepath, err := c.CreateFile(file, data)
		if err != nil {
			failed = append(failed, err.Error())
			continue
		}

		created = append(created, filepath)
	}

	if len(failed) > 0 {
		errs := ""
		for _, f := range failed {
			errs = errs + fmt.Sprintf(" - %s\n", f)
		}

		return nil, errors.New(fmt.Sprintf("Failed with errors:\n%s", errs))
	}

	return created, nil
}
