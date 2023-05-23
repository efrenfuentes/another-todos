package validator

import "regexp"

type Validator struct {
	Errors map[string]string
}

// New() initializes a new Validator instance.
func New() *Validator {
	return &Validator{
		Errors: make(map[string]string),
	}
}

// Valid() returns true if there are no errors, otherwise it returns false.
func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

// AddError() adds an error message for a given field to the validator's errors map.
// So long as the field doesn't already have an error message associated with it.
func (v *Validator) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

// Check adds an error message to the validator's errors map only if a validation is not 'ok'.
func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

// In() returns true if a given value is in a list of valid values.
func (v *Validator) In(value string, list ...string) bool {
	for _, v := range list {
		if value == v {
			return true
		}
	}
	return false
}

// Matches returns true if a string value matches a specific regexp pattern
func (v *Validator) Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

// MinLength() returns true if a string value is at least a specific length.
func (v *Validator) MinLength(value string, d int) bool {
	return len(value) >= d
}

// MaxLength() returns true if a string value is at most a specific length.
func (v *Validator) MaxLength(value string, d int) bool {
	return len(value) <= d
}

// Min() returns true if an integer value is at least a specific value.
func (v *Validator) Min(value, d int) bool {
	return value >= d
}

// Max() returns true if an integer value is at most a specific value.
func (v *Validator) Max(value, d int) bool {
	return value <= d
}

// Unique() returns true if all the strings in a slice are unique.
func (v *Validator) Unique(values []string) bool {
	uniqueValues := make(map[string]bool)

	for _, value := range values {
		uniqueValues[value] = true
	}

	return len(uniqueValues) == len(values)
}

// NoWhitespace() returns true if a string value contains no whitespace.
func (v *Validator) NoWhitespace(value string) bool {
	return !regexp.MustCompile(`\s`).MatchString(value)
}

// IsEmail() returns true if a string value is in the correct format for an email address.
func (v *Validator) IsEmail(value string) bool {
	return v.Matches(value, regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_`+"`"+`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`))
}

// IsURL() returns true if a string value is in the correct format for a URL.
func (v *Validator) IsURL(value string) bool {
	return v.Matches(value, regexp.MustCompile(`^https?://([a-z0-9-]+\.)+[a-z]{2,}(:\d{1,5})?(/.*)?$`))
}

// IsImageURL() returns true if a string value is in the correct format for a URL that points to an image file.
func (v *Validator) IsImageURL(value string) bool {
	return v.Matches(value, regexp.MustCompile(`^https?://([a-z0-9-]+\.)+[a-z]{2,}(:\d{1,5})?(/.*)?\.(jpg|jpeg|png|gif)$`))
}

// IsUUID() returns true if a string value is in the correct format for a UUID.
func (v *Validator) IsUUID(value string) bool {
	return v.Matches(value, regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`))
}
