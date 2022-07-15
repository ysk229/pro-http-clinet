// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/api/weather/v1/weather.proto

package weather

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Request with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Request with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in RequestMultiError, or nil if none found.
func (m *Request) ValidateAll() error {
	return m.validate(true)
}

func (m *Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Key

	if utf8.RuneCountInString(m.GetLocation()) != 7 {
		err := RequestValidationError{
			field:  "Location",
			reason: "value length must be 7 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if len(errors) > 0 {
		return RequestMultiError(errors)
	}

	return nil
}

// RequestMultiError is an error wrapping multiple validation errors returned
// by Request.ValidateAll() if the designated constraints aren't met.
type RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RequestMultiError) AllErrors() []error { return m }

// RequestValidationError is the validation error returned by Request.Validate
// if the designated constraints aren't met.
type RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RequestValidationError) ErrorName() string { return "RequestValidationError" }

// Error satisfies the builtin error interface
func (e RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RequestValidationError{}

// Validate checks the field values on Reply with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Reply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Reply with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ReplyMultiError, or nil if none found.
func (m *Reply) ValidateAll() error {
	return m.validate(true)
}

func (m *Reply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetResults() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ReplyValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ReplyValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ReplyValidationError{
					field:  fmt.Sprintf("Results[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ReplyMultiError(errors)
	}

	return nil
}

// ReplyMultiError is an error wrapping multiple validation errors returned by
// Reply.ValidateAll() if the designated constraints aren't met.
type ReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReplyMultiError) AllErrors() []error { return m }

// ReplyValidationError is the validation error returned by Reply.Validate if
// the designated constraints aren't met.
type ReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReplyValidationError) ErrorName() string { return "ReplyValidationError" }

// Error satisfies the builtin error interface
func (e ReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReplyValidationError{}

// Validate checks the field values on Results with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Results) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Results with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ResultsMultiError, or nil if none found.
func (m *Results) ValidateAll() error {
	return m.validate(true)
}

func (m *Results) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetLocation()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ResultsValidationError{
					field:  "Location",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ResultsValidationError{
					field:  "Location",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLocation()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ResultsValidationError{
				field:  "Location",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetDaily() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ResultsValidationError{
						field:  fmt.Sprintf("Daily[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ResultsValidationError{
						field:  fmt.Sprintf("Daily[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ResultsValidationError{
					field:  fmt.Sprintf("Daily[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for LastUpdate

	if len(errors) > 0 {
		return ResultsMultiError(errors)
	}

	return nil
}

// ResultsMultiError is an error wrapping multiple validation errors returned
// by Results.ValidateAll() if the designated constraints aren't met.
type ResultsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ResultsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ResultsMultiError) AllErrors() []error { return m }

// ResultsValidationError is the validation error returned by Results.Validate
// if the designated constraints aren't met.
type ResultsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResultsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResultsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResultsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResultsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResultsValidationError) ErrorName() string { return "ResultsValidationError" }

// Error satisfies the builtin error interface
func (e ResultsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResults.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResultsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResultsValidationError{}

// Validate checks the field values on Daily with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Daily) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Daily with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in DailyMultiError, or nil if none found.
func (m *Daily) ValidateAll() error {
	return m.validate(true)
}

func (m *Daily) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Country

	// no validation rules for Path

	// no validation rules for Timezone

	// no validation rules for TimezoneOffset

	if len(errors) > 0 {
		return DailyMultiError(errors)
	}

	return nil
}

// DailyMultiError is an error wrapping multiple validation errors returned by
// Daily.ValidateAll() if the designated constraints aren't met.
type DailyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DailyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DailyMultiError) AllErrors() []error { return m }

// DailyValidationError is the validation error returned by Daily.Validate if
// the designated constraints aren't met.
type DailyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DailyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DailyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DailyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DailyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DailyValidationError) ErrorName() string { return "DailyValidationError" }

// Error satisfies the builtin error interface
func (e DailyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDaily.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DailyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DailyValidationError{}

// Validate checks the field values on Location with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Location) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Location with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LocationMultiError, or nil
// if none found.
func (m *Location) ValidateAll() error {
	return m.validate(true)
}

func (m *Location) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Country

	// no validation rules for Path

	// no validation rules for Timezone

	// no validation rules for TimezoneOffset

	if len(errors) > 0 {
		return LocationMultiError(errors)
	}

	return nil
}

// LocationMultiError is an error wrapping multiple validation errors returned
// by Location.ValidateAll() if the designated constraints aren't met.
type LocationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LocationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LocationMultiError) AllErrors() []error { return m }

// LocationValidationError is the validation error returned by
// Location.Validate if the designated constraints aren't met.
type LocationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LocationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LocationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LocationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LocationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LocationValidationError) ErrorName() string { return "LocationValidationError" }

// Error satisfies the builtin error interface
func (e LocationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLocation.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LocationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LocationValidationError{}
