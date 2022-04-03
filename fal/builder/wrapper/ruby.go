package wrapper

import (
	_ "embed"
	"fal/shared/fs"
)

var (
	//go:embed ruby/main.rb
	_ruby_main []byte

	//go:embed ruby/Gemfile
	_ruby_gem []byte

	//go:embed ruby/Gemfile.lock
	_ruby_lock []byte
)

func BuildRubyWrapper(wd *fs.Location) error {
	ld := wd.InnerLevel("ruby")

	err := ld.CreateDir()
	if err != nil {
		return err
	}

	f := fs.FileList{
		"main.rb":      _ruby_main,
		"Gemfile":      _ruby_gem,
		"Gemfile.lock": _ruby_lock,
	}

	_, err = ld.CreateManyFiles(f)
	if err != nil {
		return err
	}

	return nil
}
