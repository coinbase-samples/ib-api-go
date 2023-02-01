// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pkg/pbs/asset/v1/asset.proto

package v1

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

// Validate checks the field values on ListAssetsRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListAssetsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListAssetsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListAssetsRequestMultiError, or nil if none found.
func (m *ListAssetsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListAssetsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListAssetsRequestMultiError(errors)
	}

	return nil
}

// ListAssetsRequestMultiError is an error wrapping multiple validation errors
// returned by ListAssetsRequest.ValidateAll() if the designated constraints
// aren't met.
type ListAssetsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListAssetsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListAssetsRequestMultiError) AllErrors() []error { return m }

// ListAssetsRequestValidationError is the validation error returned by
// ListAssetsRequest.Validate if the designated constraints aren't met.
type ListAssetsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListAssetsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListAssetsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListAssetsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListAssetsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListAssetsRequestValidationError) ErrorName() string {
	return "ListAssetsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListAssetsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListAssetsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListAssetsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListAssetsRequestValidationError{}

// Validate checks the field values on ListAssetsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListAssetsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListAssetsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListAssetsResponseMultiError, or nil if none found.
func (m *ListAssetsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListAssetsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetData() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListAssetsResponseValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListAssetsResponseValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListAssetsResponseValidationError{
					field:  fmt.Sprintf("Data[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListAssetsResponseMultiError(errors)
	}

	return nil
}

// ListAssetsResponseMultiError is an error wrapping multiple validation errors
// returned by ListAssetsResponse.ValidateAll() if the designated constraints
// aren't met.
type ListAssetsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListAssetsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListAssetsResponseMultiError) AllErrors() []error { return m }

// ListAssetsResponseValidationError is the validation error returned by
// ListAssetsResponse.Validate if the designated constraints aren't met.
type ListAssetsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListAssetsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListAssetsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListAssetsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListAssetsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListAssetsResponseValidationError) ErrorName() string {
	return "ListAssetsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListAssetsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListAssetsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListAssetsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListAssetsResponseValidationError{}

// Validate checks the field values on Asset with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Asset) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Asset with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in AssetMultiError, or nil if none found.
func (m *Asset) ValidateAll() error {
	return m.validate(true)
}

func (m *Asset) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AssetId

	// no validation rules for Ticker

	// no validation rules for Name

	// no validation rules for MinTransactionAmount

	// no validation rules for MaxTransactionAmount

	// no validation rules for HighOffer

	// no validation rules for LowBid

	// no validation rules for Slippage

	// no validation rules for Spread

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AssetValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AssetValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AssetValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for MarketCap

	// no validation rules for Volume

	// no validation rules for Supply

	// no validation rules for Direction

	if len(errors) > 0 {
		return AssetMultiError(errors)
	}

	return nil
}

// AssetMultiError is an error wrapping multiple validation errors returned by
// Asset.ValidateAll() if the designated constraints aren't met.
type AssetMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AssetMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AssetMultiError) AllErrors() []error { return m }

// AssetValidationError is the validation error returned by Asset.Validate if
// the designated constraints aren't met.
type AssetValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AssetValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AssetValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AssetValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AssetValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AssetValidationError) ErrorName() string { return "AssetValidationError" }

// Error satisfies the builtin error interface
func (e AssetValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAsset.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AssetValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AssetValidationError{}

// Validate checks the field values on GetAssetRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetAssetRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetAssetRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetAssetRequestMultiError, or nil if none found.
func (m *GetAssetRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetAssetRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetId()) != 36 {
		err := GetAssetRequestValidationError{
			field:  "Id",
			reason: "value length must be 36 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if len(errors) > 0 {
		return GetAssetRequestMultiError(errors)
	}

	return nil
}

// GetAssetRequestMultiError is an error wrapping multiple validation errors
// returned by GetAssetRequest.ValidateAll() if the designated constraints
// aren't met.
type GetAssetRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetAssetRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetAssetRequestMultiError) AllErrors() []error { return m }

// GetAssetRequestValidationError is the validation error returned by
// GetAssetRequest.Validate if the designated constraints aren't met.
type GetAssetRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAssetRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAssetRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAssetRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAssetRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAssetRequestValidationError) ErrorName() string { return "GetAssetRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetAssetRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAssetRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAssetRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAssetRequestValidationError{}

// Validate checks the field values on GetAssetResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetAssetResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetAssetResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetAssetResponseMultiError, or nil if none found.
func (m *GetAssetResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetAssetResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetAssetResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetAssetResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetAssetResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetAssetResponseMultiError(errors)
	}

	return nil
}

// GetAssetResponseMultiError is an error wrapping multiple validation errors
// returned by GetAssetResponse.ValidateAll() if the designated constraints
// aren't met.
type GetAssetResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetAssetResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetAssetResponseMultiError) AllErrors() []error { return m }

// GetAssetResponseValidationError is the validation error returned by
// GetAssetResponse.Validate if the designated constraints aren't met.
type GetAssetResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAssetResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAssetResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAssetResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAssetResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAssetResponseValidationError) ErrorName() string { return "GetAssetResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetAssetResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAssetResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAssetResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAssetResponseValidationError{}
