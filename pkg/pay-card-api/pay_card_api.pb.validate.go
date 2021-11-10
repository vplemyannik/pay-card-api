// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: ozonmp/pay_card_api/v1/pay_card_api.proto

package pay_card_api

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

// Validate checks the field values on Card with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Card) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetOwnerId() <= 0 {
		return CardValidationError{
			field:  "OwnerId",
			reason: "value must be greater than 0",
		}
	}

	if utf8.RuneCountInString(m.GetPaymentSystem()) < 3 {
		return CardValidationError{
			field:  "PaymentSystem",
			reason: "value length must be at least 3 runes",
		}
	}

	if utf8.RuneCountInString(m.GetNumber()) != 16 {
		return CardValidationError{
			field:  "Number",
			reason: "value length must be 16 runes",
		}

	}

	if utf8.RuneCountInString(m.GetHolderName()) < 2 {
		return CardValidationError{
			field:  "HolderName",
			reason: "value length must be at least 2 runes",
		}
	}

	if utf8.RuneCountInString(m.GetCvcCvv()) != 3 {
		return CardValidationError{
			field:  "CvcCvv",
			reason: "value length must be 3 runes",
		}

	}

	if m.GetExpirationDate() == nil {
		return CardValidationError{
			field:  "ExpirationDate",
			reason: "value is required",
		}
	}

	return nil
}

// CardValidationError is the validation error returned by Card.Validate if the
// designated constraints aren't met.
type CardValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CardValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CardValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CardValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CardValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CardValidationError) ErrorName() string { return "CardValidationError" }

// Error satisfies the builtin error interface
func (e CardValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCard.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CardValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CardValidationError{}

// Validate checks the field values on RemoveCardV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveCardV1Request) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// RemoveCardV1RequestValidationError is the validation error returned by
// RemoveCardV1Request.Validate if the designated constraints aren't met.
type RemoveCardV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveCardV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveCardV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveCardV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveCardV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveCardV1RequestValidationError) ErrorName() string {
	return "RemoveCardV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveCardV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveCardV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveCardV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveCardV1RequestValidationError{}

// Validate checks the field values on ListCardV1Request with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListCardV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetOffset() <= 0 {
		return ListCardV1RequestValidationError{
			field:  "Offset",
			reason: "value must be greater than 0",
		}
	}

	if val := m.GetLimit(); val < 0 || val > 100 {
		return ListCardV1RequestValidationError{
			field:  "Limit",
			reason: "value must be inside range [0, 100]",
		}
	}

	return nil
}

// ListCardV1RequestValidationError is the validation error returned by
// ListCardV1Request.Validate if the designated constraints aren't met.
type ListCardV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListCardV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListCardV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListCardV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListCardV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListCardV1RequestValidationError) ErrorName() string {
	return "ListCardV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListCardV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListCardV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListCardV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListCardV1RequestValidationError{}

// Validate checks the field values on DescribeCardV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeCardV1Request) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// DescribeCardV1RequestValidationError is the validation error returned by
// DescribeCardV1Request.Validate if the designated constraints aren't met.
type DescribeCardV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeCardV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeCardV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeCardV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeCardV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeCardV1RequestValidationError) ErrorName() string {
	return "DescribeCardV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeCardV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeCardV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeCardV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeCardV1RequestValidationError{}

// Validate checks the field values on CreateCardV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateCardV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// CreateCardV1ResponseValidationError is the validation error returned by
// CreateCardV1Response.Validate if the designated constraints aren't met.
type CreateCardV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateCardV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateCardV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateCardV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateCardV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateCardV1ResponseValidationError) ErrorName() string {
	return "CreateCardV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateCardV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateCardV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateCardV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateCardV1ResponseValidationError{}

// Validate checks the field values on ListCardV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListCardV1Response) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetCards() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListCardV1ResponseValidationError{
					field:  fmt.Sprintf("Cards[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListCardV1ResponseValidationError is the validation error returned by
// ListCardV1Response.Validate if the designated constraints aren't met.
type ListCardV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListCardV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListCardV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListCardV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListCardV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListCardV1ResponseValidationError) ErrorName() string {
	return "ListCardV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListCardV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListCardV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListCardV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListCardV1ResponseValidationError{}
