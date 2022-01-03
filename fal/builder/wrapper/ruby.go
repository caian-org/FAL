package wrapper

import (
	_ "embed"
	"fal/util"
)

var (
	//go:embed ruby/fal.rb
	_ruby_main []byte

	//go:embed ruby/Gemfile
	_ruby_gem []byte

	//go:embed ruby/Gemfile.lock
	_ruby_lock []byte
)

func BuildRubyWrapper(wd *util.Location) error {
	ld := wd.InnerLevel("ruby")

	err := ld.CreateDir()
	if err != nil {
		return err
	}

	f := util.FileList{
		"Gemfile":      _ruby_gem,
		"Gemfile.lock": _ruby_lock,
	}

	_, err = ld.CreateManyFiles(f)
	if err != nil {
		return err
	}

	return nil
}
