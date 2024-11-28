// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: web/v1/common.proto

package web

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

// Validate checks the field values on CommonSendSmsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CommonSendSmsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CommonSendSmsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CommonSendSmsRequestMultiError, or nil if none found.
func (m *CommonSendSmsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CommonSendSmsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Mobile

	// no validation rules for Channel

	if len(errors) > 0 {
		return CommonSendSmsRequestMultiError(errors)
	}

	return nil
}

// CommonSendSmsRequestMultiError is an error wrapping multiple validation
// errors returned by CommonSendSmsRequest.ValidateAll() if the designated
// constraints aren't met.
type CommonSendSmsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CommonSendSmsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CommonSendSmsRequestMultiError) AllErrors() []error { return m }

// CommonSendSmsRequestValidationError is the validation error returned by
// CommonSendSmsRequest.Validate if the designated constraints aren't met.
type CommonSendSmsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommonSendSmsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommonSendSmsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommonSendSmsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommonSendSmsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommonSendSmsRequestValidationError) ErrorName() string {
	return "CommonSendSmsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CommonSendSmsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommonSendSmsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommonSendSmsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommonSendSmsRequestValidationError{}

// Validate checks the field values on CommonSendSmsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CommonSendSmsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CommonSendSmsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CommonSendSmsResponseMultiError, or nil if none found.
func (m *CommonSendSmsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CommonSendSmsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CommonSendSmsResponseMultiError(errors)
	}

	return nil
}

// CommonSendSmsResponseMultiError is an error wrapping multiple validation
// errors returned by CommonSendSmsResponse.ValidateAll() if the designated
// constraints aren't met.
type CommonSendSmsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CommonSendSmsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CommonSendSmsResponseMultiError) AllErrors() []error { return m }

// CommonSendSmsResponseValidationError is the validation error returned by
// CommonSendSmsResponse.Validate if the designated constraints aren't met.
type CommonSendSmsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommonSendSmsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommonSendSmsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommonSendSmsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommonSendSmsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommonSendSmsResponseValidationError) ErrorName() string {
	return "CommonSendSmsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CommonSendSmsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommonSendSmsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommonSendSmsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommonSendSmsResponseValidationError{}

// Validate checks the field values on CommonSendEmailRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CommonSendEmailRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CommonSendEmailRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CommonSendEmailRequestMultiError, or nil if none found.
func (m *CommonSendEmailRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CommonSendEmailRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Email

	if len(errors) > 0 {
		return CommonSendEmailRequestMultiError(errors)
	}

	return nil
}

// CommonSendEmailRequestMultiError is an error wrapping multiple validation
// errors returned by CommonSendEmailRequest.ValidateAll() if the designated
// constraints aren't met.
type CommonSendEmailRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CommonSendEmailRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CommonSendEmailRequestMultiError) AllErrors() []error { return m }

// CommonSendEmailRequestValidationError is the validation error returned by
// CommonSendEmailRequest.Validate if the designated constraints aren't met.
type CommonSendEmailRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommonSendEmailRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommonSendEmailRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommonSendEmailRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommonSendEmailRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommonSendEmailRequestValidationError) ErrorName() string {
	return "CommonSendEmailRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CommonSendEmailRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommonSendEmailRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommonSendEmailRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommonSendEmailRequestValidationError{}

// Validate checks the field values on CommonSendEmailResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CommonSendEmailResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CommonSendEmailResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CommonSendEmailResponseMultiError, or nil if none found.
func (m *CommonSendEmailResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CommonSendEmailResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CommonSendEmailResponseMultiError(errors)
	}

	return nil
}

// CommonSendEmailResponseMultiError is an error wrapping multiple validation
// errors returned by CommonSendEmailResponse.ValidateAll() if the designated
// constraints aren't met.
type CommonSendEmailResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CommonSendEmailResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CommonSendEmailResponseMultiError) AllErrors() []error { return m }

// CommonSendEmailResponseValidationError is the validation error returned by
// CommonSendEmailResponse.Validate if the designated constraints aren't met.
type CommonSendEmailResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommonSendEmailResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommonSendEmailResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommonSendEmailResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommonSendEmailResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommonSendEmailResponseValidationError) ErrorName() string {
	return "CommonSendEmailResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CommonSendEmailResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommonSendEmailResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommonSendEmailResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommonSendEmailResponseValidationError{}
