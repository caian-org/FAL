package wrapper

import (
	"errors"
	"fmt"

	"fal/shared/fs"
	. "fal/wrapper/target"
)

type WrapperBuilderFunc func(location *fs.Location) error

type WrapperLang = int32

const (
	WrapperRuby WrapperLang = iota
	WrapperPython
	WrapperJavascript
)

func targetFieldToWrapperLang(target string) (WrapperLang, error) {
	if target == "ruby" {
		return WrapperRuby, nil
	}

	if target == "python" {
		return WrapperPython, nil
	}

	if target == "javascript" {
		return WrapperJavascript, nil
	}

	return -1, errors.New(fmt.Sprintf("Unsupported target '%s'", target))
}

func GetWrapperBuilderOf(target string) (WrapperBuilderFunc, error) {
	lang, err := targetFieldToWrapperLang(target)
	if err != nil {
		return nil, err
	}

	switch lang {
	case WrapperRuby:
		return WrapperRubyBuilder, nil

	case WrapperPython:
		return WrapperPythonBuilder, nil

	case WrapperJavascript:
		return WrapperJavascriptBuilder, nil
	}

	return nil, errors.New(fmt.Sprintf("Target '%s' is not implemented yet", target))
}
