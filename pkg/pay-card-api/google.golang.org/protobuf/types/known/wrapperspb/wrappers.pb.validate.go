// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: google/protobuf/wrappers.proto

package wrapperspb

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
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
)

// Validate checks the field values on DoubleValue with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *DoubleValue) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Value

	return nil
}

// DoubleValueValidationError is the validation error returned by
// DoubleValue.Validate if the designated constraints aren't met.
type DoubleValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DoubleValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DoubleValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DoubleValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DoubleValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DoubleValueValidationError) ErrorName() string { return "DoubleValueValidationError" }

// Error satisfies the builtin error interface
func (e DoubleValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDoubleValue.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DoubleValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DoubleValueValidationError{}

// Validate checks the field values on FloatValue with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *FloatValue) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Value

	return nil
}

// FloatValueValidationError is the validation error returned by
// FloatValue.Validate if the designated constraints aren't met.
type FloatValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FloatValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FloatValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FloatValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FloatValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FloatValueValidationError) ErrorName() string { return "FloatValueValidationError" }

// Error satisfies the builtin error interface
func (e FloatValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFloatValue.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FloatValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FloatValueValidationError{}

// Validate checks the field values on Int64Value with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Int64Value) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Value

	return nil
}

// Int64ValueValidationError is the validation error returned by
// Int64Value.Validate if the designated constraints aren't met.
type Int64ValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Int64ValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Int64ValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Int64ValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Int64ValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Int64ValueValidationError) ErrorName() string { return "Int64ValueValidationError" }

// Error satisfies the builtin error interface
func (e Int64ValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInt64Value.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Int64ValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Int64ValueValidationError{}

// Validate checks the field values on UInt64Value with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *UInt64Value) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Value

	return nil
}

// UInt64ValueValidationError is the validation error returned by
// UInt64Value.Validate if the designated constraints aren't met.
type UInt64ValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UInt64ValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UInt64ValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UInt64ValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UInt64ValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UInt64ValueValidationError) ErrorName() string { return "UInt64ValueValidationError" }

// Error satisfies the builtin error interface
func (e UInt64ValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUInt64Value.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UInt64ValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UInt64ValueValidationError{}

// Validate checks the field values on Int32Value with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Int32Value) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Value

	return nil
}

// Int32ValueValidationError is the validation error returned by
// Int32Value.Validate if the designated constraints aren't met.
type Int32ValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Int32ValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Int32ValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Int32ValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Int32ValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Int32ValueValidationError) ErrorName() string { return "Int32ValueValidationError" }

// Error satisfies the builtin error interface
func (e Int32ValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInt32Value.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Int32ValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Int32ValueValidationError{}

// Validate checks the field values on UInt32Value with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *UInt32Value) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Value

	return nil
}

// UInt32ValueValidationError is the validation error returned by
// UInt32Value.Validate if the designated constraints aren't met.
type UInt32ValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UInt32ValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UInt32ValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UInt32ValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UInt32ValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UInt32ValueValidationError) ErrorName() string { return "UInt32ValueValidationError" }

// Error satisfies the builtin error interface
func (e UInt32ValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUInt32Value.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UInt32ValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UInt32ValueValidationError{}

// Validate checks the field values on BoolValue with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *BoolValue) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Value

	return nil
}

// BoolValueValidationError is the validation error returned by
// BoolValue.Validate if the designated constraints aren't met.
type BoolValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BoolValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BoolValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BoolValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BoolValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BoolValueValidationError) ErrorName() string { return "BoolValueValidationError" }

// Error satisfies the builtin error interface
func (e BoolValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBoolValue.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BoolValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BoolValueValidationError{}

// Validate checks the field values on StringValue with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *StringValue) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Value

	return nil
}

// StringValueValidationError is the validation error returned by
// StringValue.Validate if the designated constraints aren't met.
type StringValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StringValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StringValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StringValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StringValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StringValueValidationError) ErrorName() string { return "StringValueValidationError" }

// Error satisfies the builtin error interface
func (e StringValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStringValue.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StringValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StringValueValidationError{}

// Validate checks the field values on BytesValue with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *BytesValue) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Value

	return nil
}

// BytesValueValidationError is the validation error returned by
// BytesValue.Validate if the designated constraints aren't met.
type BytesValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BytesValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BytesValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BytesValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BytesValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BytesValueValidationError) ErrorName() string { return "BytesValueValidationError" }

// Error satisfies the builtin error interface
func (e BytesValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBytesValue.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BytesValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BytesValueValidationError{}
