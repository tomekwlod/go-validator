package validator

import (
	"fmt"
	"reflect"

	"github.com/tomekwlod/go-validator"
)

type OptFunc func(*Opts)

type Opts struct {
	message string
}

func defaultOpts() Opts {
	return Opts{
		message: "The provided value must be bigger or equal %d",
	}
}

func WithMessage(s string) OptFunc {
	return func(opts *Opts) {
		opts.message = s
	}
}

type minValidator struct {
	Opts
	min int64
}

// TODO: generic please
func New(min int64, opts ...OptFunc) *minValidator {
	o := defaultOpts()

	for _, fn := range opts {
		fn(&o)
	}

	return &minValidator{
		min:  min,
		Opts: o,
	}
}

func (n minValidator) Validate(v any) (bool, string) {
	var _value any

	// Check if v is nil
	if validator.IsNil(v) {
		return false, fmt.Sprintf(n.message, n.min)
	}

	// Check if v is a pointer
	if reflect.TypeOf(v).Kind() == reflect.Ptr {
		// Dereference the pointer
		val := reflect.ValueOf(v).Elem()

		// Get the underlying value
		_value = val.Interface()
	} else {
		_value = v
	}

	// Handle different types
	switch value := _value.(type) {
	case int, int64, float64, uint, uint64, uint32, uint16, uint8, int32, int16, int8:
		if value.(int64) < n.min {
			return false, fmt.Sprintf(n.message, n.min)
		} else {
			return true, ""
		}

	default:
		return false, fmt.Sprintf("unsupported type: %T", _value)
	}
}
