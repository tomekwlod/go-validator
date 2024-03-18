package validator

import "reflect"

// Function to check if a value is nil using reflection
func IsNil(v interface{}) bool {
	if v == nil {
		return true
	}

	value := reflect.ValueOf(v)
	kind := value.Kind()

	if kind == reflect.Ptr || kind == reflect.Interface {
		return value.IsNil()
	}

	// Using reflection to check for zero value
	zero := reflect.Zero(reflect.TypeOf(value)).Interface()
	return reflect.DeepEqual(value, zero)
}
