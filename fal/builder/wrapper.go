package builder

import (
	"errors"
	"fmt"

	"fal/builder/wrapper"
	"fal/util"
)

type WrapperBuilderFunc func(location *util.Location) error

type WrapperLang = int32

const (
	WrapperRuby WrapperLang = iota
	WrapperPython
	WrapperJavaScript
)

func targetFieldToWrapperLang(target string) (WrapperLang, error) {
	if target == "ruby" {
		return WrapperRuby, nil
	}

	if target == "python" {
		return WrapperPython, nil
	}

	if target == "javascript" {
		return WrapperJavaScript, nil
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
		return wrapper.BuildRubyWrapper, nil

	case WrapperPython:
		return wrapper.BuildPythonWrapper, nil

	case WrapperJavaScript:
		return wrapper.BuildJavaScriptWrapper, nil
	}

	return nil, errors.New(fmt.Sprintf("Target '%s' is not implemented yet", target))
}
