package validator

import "github.com/asaskevich/govalidator"

type OptFunc func(*Opts)

type Opts struct {
	notValidMessage string
	charMessage     string
	message         string
}

func defaultOpts() Opts {
	return Opts{
		notValidMessage: "Please enter a valid email address.",
		charMessage:     "Your email address contains invalid characters. Please check your details and enter a valid email address.",
		message:         "", // this overrides all the above messages
	}
}

//	func withTLS(opts *Opts) {
//		fmt.Println("tls exec")
//		opts.tls = true
//	}
func WithMessage(s string) OptFunc {
	return func(opts *Opts) {
		opts.notValidMessage = s
		opts.charMessage = s
		opts.message = s
	}
}
func WithNotValidMessage(s string) OptFunc {
	return func(opts *Opts) {
		opts.notValidMessage = s
	}
}
func WithCharMessage(s string) OptFunc {
	return func(opts *Opts) {
		opts.charMessage = s
	}
}

type emailValidator struct {
	Opts
}

func New(opts ...OptFunc) *emailValidator {
	o := defaultOpts()

	for _, fn := range opts {
		fn(&o)
	}

	return &emailValidator{
		Opts: o,
	}
}

func (n emailValidator) Validate(v any) (bool, string) {
	str := v.(string)

	switch {
	case !govalidator.IsEmail(str):
		return false, n.notValidMessage
	case !govalidator.Matches(str, "^[a-zA-Z0-9@\\-\\._\\=,\\+]+$"):
		return false, n.charMessage
	}

	return true, ""
}
