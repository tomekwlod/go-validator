package validator

import (
	"reflect"

	gv "github.com/tomekwlod/go-validator"
)

type OptFunc func(*Opts)

type Opts struct {
	message string
}

func defaultOpts() Opts {
	return Opts{
		message: "This value cannot be empty",
	}
}

func WithMessage(s string) OptFunc {
	return func(opts *Opts) {
		opts.message = s
	}
}

type neValidator struct {
	Opts
}

func New(opts ...OptFunc) *neValidator {
	o := defaultOpts()

	for _, fn := range opts {
		fn(&o)
	}

	return &neValidator{
		Opts: o,
	}
}

func (n neValidator) Validate(v any) (bool, string) {

	// Check if v is nil
	if gv.IsNil(v) {
		return false, n.message
	}

	// Check if v is a pointer
	if reflect.TypeOf(v).Kind() == reflect.Ptr {
		// Dereference the pointer
		val := reflect.ValueOf(v).Elem()

		// Get the underlying value
		value := val.Interface()

		// Handle different types
		switch value := value.(type) {
		// case int:
		// 	// if value == 0 {
		// 	// 	return false, n.message
		// 	// } else {
		// 	return false, n.message
		// 	// }
		// case float64:
		// 	// if value == 0 {
		// 	// 	fmt.Println("Value is zero (float64)")
		// 	// } else {
		// 	return false, n.message
		// 	// }
		case []*int:
		case []*int64:
		case []int:
		case []int64:
		case []string:
			if len(value) == 0 {
				return false, n.message
			}
		case string:
			if value == "" {
				return false, n.message
			}
			// default:
			// 	fmt.Println("Unsupported type:", reflect.TypeOf(value))
		}
	} else {
		// Handle non-pointer types
		switch value := v.(type) {
		// case int:
		// if value == 0 {
		// 	fmt.Println("Value is zero (int)")
		// } else {
		// return false, n.message
		// }
		// case float64:
		// if value == 0 {
		// 	fmt.Println("Value is zero (float64)")
		// } else {
		// return false, n.message
		// }

		// TODO: the isNil should be already doing this... we need tests with all possible cases and adjust the isNil function
		case []*int:
			if len(value) == 0 {
				return false, n.message
			}
		// TODO: the isNil should be already doing this... we need tests with all possible cases and adjust the isNil function
		case []int:
			if len(value) == 0 {
				return false, n.message
			}

		case string:
			if value == "" {
				return false, n.message
			}
			// } else {
			// 	fmt.Println("Value:", value)
			// }
			// default:
			// 	fmt.Println("Unsupported type:", reflect.TypeOf(value))
		}
	}

	return true, ""
}
