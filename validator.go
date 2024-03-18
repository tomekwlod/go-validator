package validator

import (
	"net/url"
)

type Validatorer interface {
	Validate(v any) (bool, string)
}

type validator struct {
	fields []func(*url.Values)
}

func New() *validator {
	return &validator{}
}

func (v *validator) Add(name string, value any, vs ...Validatorer) {
	for _, _v := range vs {
		// https://yourbasic.org/golang/gotcha-range-copy-array/
		// https://medium.com/@nsspathirana/common-mistakes-with-go-slices-95f2e9b362a9#:~:text=to%20that%20variable.-,Example,-with%20the%20mistake
		_tmp := _v

		v.fields = append(v.fields, func(e *url.Values) {
			if ok, errmsg := _tmp.Validate(value); !ok {
				e.Add(name, errmsg)
			}
		})

	}
}

func (v *validator) Validate() url.Values {
	errs := url.Values{}

	for _, f := range v.fields {
		f(&errs)
	}

	return errs
}
