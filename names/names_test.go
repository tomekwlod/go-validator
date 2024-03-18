//go:build unit

package validator

import (
	"testing"
)

func TestNames(t *testing.T) {
	nv := NewNamesValidator(NamesParams{Max: 90, CharMessage: "CharError"})

	names := map[string]string{
		"łôöòóœøōõàáâäæãåāßśšèéêëēėęûüùúūïìžźżçćčñń":  "",
		"łôöòóœøōõàáâäæãåāßśš èéêëēėęûüùúūïìžźżçćčñń": "",
		"test":              "",
		" test":             "CharError",
		" test ":            "CharError",
		"test ":             "CharError",
		"=+$5£@":            "CharError",
		"Diablo-Włodarczyk": "",
		"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890": "The provided value is too long",
	}

	iteration := 1

	for n, e := range names {
		iteration++

		ok, got := nv.Validate(s)

		if got == "" {
			if !ok {
				t.Errorf("[%d] for string `%s` expected test to pass, got: `%s`", iteration, n, e)
			}
		} else {
			if got != e {
				t.Errorf("[%d] for string `%s` expected test to FAIL with a message `%s`, got: `%s`", iteration, n, e, got)
			}
		}
	}
}
