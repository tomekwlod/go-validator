package validator

import "unicode"

type OptFunc func(*Opts)

type Opts struct {
	message string
}

func defaultOpts() Opts {
	return Opts{
		message: "Your password is not strong enough. Your password must be at least 7 characters long and contain at least one capital letter, at least one number and at least one special character (e.g. ! = * £).",
	}
}

func WithMessage(s string) OptFunc {
	return func(opts *Opts) {
		opts.message = s
	}
}

type passwordValidator struct {
	Opts
}

func New(opts ...OptFunc) *passwordValidator {
	o := defaultOpts()

	for _, fn := range opts {
		fn(&o)
	}

	return &passwordValidator{
		Opts: o,
	}
}

func (n passwordValidator) Validate(v any) (bool, string) {
	str := v.(string)

	var letters int

	var number, special, sevenOrMore, upper bool

	for _, c := range str {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
			// letters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		// case unicode.IsLetter(c) || c == ' ':
		// 	letters++
		default:
			//return false, false, false, false
		}
		letters++
	}

	sevenOrMore = letters >= 7

	if !sevenOrMore || !special || !upper || !number {
		return false, n.message
	}

	return true, ""
}
