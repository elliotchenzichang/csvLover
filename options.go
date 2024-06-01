package csvLover

import "errors"

type Options struct {
	Path       string
	Comma      rune
	LazyQuotes bool
	Dest       any
}

func (opt *Options) validate() error {
	if opt == nil {
		return errors.New("nil options")
	}
	if opt.Path == "" {
		return errors.New("path is required")
	}
	return nil
}
