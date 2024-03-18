package validator

import "github.com/asaskevich/govalidator"

type OptFunc func(*Opts)

type Opts struct {
	min         int
	max         int
	charMessage string
	minMessage  string
	maxMessage  string
	message     string
}

func defaultOpts() Opts {
	return Opts{
		min:         1,
		max:         50,
		charMessage: "The provided value isn't valid",
		minMessage:  "The provided value is too short",
		maxMessage:  "The provided value is too long",
		message:     "", // if this is set it will override the above
	}
}

func WithMessage(s string) OptFunc {
	return func(opts *Opts) {
		opts.maxMessage = s
		opts.minMessage = s
		opts.charMessage = s
		opts.message = s
	}
}
func WithCharMessage(s string) OptFunc {
	return func(opts *Opts) {
		opts.charMessage = s
	}
}
func WithMinMessage(s string) OptFunc {
	return func(opts *Opts) {
		opts.minMessage = s
	}
}
func WithMaxMessage(s string) OptFunc {
	return func(opts *Opts) {
		opts.maxMessage = s
	}
}
func WithMin(i int) OptFunc {
	return func(opts *Opts) {
		opts.min = i
	}
}
func WithMax(i int) OptFunc {
	return func(opts *Opts) {
		opts.max = i
	}
}

type namesValidator struct {
	Opts
}

func New(opts ...OptFunc) *namesValidator {
	o := defaultOpts()

	for _, fn := range opts {
		fn(&o)
	}

	return &namesValidator{
		Opts: o,
	}
}

func (n namesValidator) Validate(v any) (bool, string) {
	str := v.(string)

	if len(str) < n.min {
		return false, n.minMessage
	}
	if len(str) > n.max {
		return false, n.maxMessage
	}
	// https://regex101.com/r/A98zgd/1
	if !govalidator.Matches(v.(string), `^[\p{L}\p{N}_][\p{L}\p{N}\p{M}_\- ']*$`) {
		return false, n.charMessage
	}

	return true, ""
}
