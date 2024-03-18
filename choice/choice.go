package validator

import (
	"fmt"
	"strings"

	"golang.org/x/exp/slices"
)

type OptFunc func(*Opts)

type Opts struct {
	message string
}

func defaultOpts(options []string) Opts {
	return Opts{
		message: fmt.Sprintf("the type value is not valid, it should be one of the following: %s", strings.Join(options, ",")),
	}
}

func WithMessage(s string) OptFunc {
	return func(opts *Opts) {
		opts.message = s
	}
}

type validator struct {
	Opts
	options []string
}

// // TODO: add generics to be able validate strings/integers
// // TODO: add Choice.options []string and create New(options) constructor --- of course not []string but generic
func New(options []string, opts ...OptFunc) *validator {
	o := defaultOpts(options)

	for _, fn := range opts {
		fn(&o)
	}

	return &validator{
		options: options,
		Opts:    o,
	}
}

func (n validator) Validate(v any) (bool, string) {
	s := v.(string)

	if slices.IndexFunc(n.options, func(o string) bool { return o == s }) < 0 {
		// IndexFunc returns the first index i satisfying f(s[i]), or -1 if none do.
		return false, n.message
	}

	return true, ""
}
