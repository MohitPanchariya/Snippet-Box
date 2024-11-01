package validator

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

// Regular expression to match email
var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// Validator is used to validate form data
type Validator struct {
	FieldErrors map[string]string
}

// Valid checks if any errors were encountered while parsing the form data
func (v *Validator) Valid() bool {
	return len(v.FieldErrors) == 0
}

// Adds an entry to v.FieldErrors if an entry for key doesn't exist already.
func (v *Validator) AddFieldError(key, message string) {
	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string]string)
	}

	if _, ok := v.FieldErrors[key]; !ok {
		v.FieldErrors[key] = message
	}
}

// Adds an entry to v.FieldErrors only if ok == false
func (v *Validator) CheckField(ok bool, key, message string) {
	if !ok {
		v.AddFieldError(key, message)
	}
}

// Returns true if value != ""
func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

// Returns true if value contains no more than n characters
func MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

// Returns true if value contains more than n characters
func MinChars(value string, n int) bool {
	return utf8.RuneCountInString(value) >= n
}

// Returns true if value matches the compiled regular expression pattern
func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

// Returns true if value is in the list of permitted values
func PermittedInt(value int, permittedValues ...int) bool {
	for i := range permittedValues {
		if value == permittedValues[i] {
			return true
		}
	}
	return false
}
