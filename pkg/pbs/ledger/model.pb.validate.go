// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: protos/ledger/model.proto

package ledger

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

// Validate checks the field values on Account with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Account) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Account with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in AccountMultiError, or nil if none found.
func (m *Account) ValidateAll() error {
	return m.validate(true)
}

func (m *Account) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetId()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AccountValidationError{
					field:  "Id",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AccountValidationError{
					field:  "Id",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccountValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetPortfolioId()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AccountValidationError{
					field:  "PortfolioId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AccountValidationError{
					field:  "PortfolioId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPortfolioId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccountValidationError{
				field:  "PortfolioId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for UserId

	// no validation rules for Currency

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AccountValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AccountValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccountValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return AccountMultiError(errors)
	}

	return nil
}

// AccountMultiError is an error wrapping multiple validation errors returned
// by Account.ValidateAll() if the designated constraints aren't met.
type AccountMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AccountMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AccountMultiError) AllErrors() []error { return m }

// AccountValidationError is the validation error returned by Account.Validate
// if the designated constraints aren't met.
type AccountValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AccountValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AccountValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AccountValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AccountValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AccountValidationError) ErrorName() string { return "AccountValidationError" }

// Error satisfies the builtin error interface
func (e AccountValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAccount.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AccountValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AccountValidationError{}

// Validate checks the field values on AccountBalance with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AccountBalance) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AccountBalance with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AccountBalanceMultiError,
// or nil if none found.
func (m *AccountBalance) ValidateAll() error {
	return m.validate(true)
}

func (m *AccountBalance) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetId()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AccountBalanceValidationError{
					field:  "Id",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AccountBalanceValidationError{
					field:  "Id",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccountBalanceValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Balance

	// no validation rules for Hold

	// no validation rules for Available

	if all {
		switch v := interface{}(m.GetAccountId()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AccountBalanceValidationError{
					field:  "AccountId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AccountBalanceValidationError{
					field:  "AccountId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAccountId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccountBalanceValidationError{
				field:  "AccountId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return AccountBalanceMultiError(errors)
	}

	return nil
}

// AccountBalanceMultiError is an error wrapping multiple validation errors
// returned by AccountBalance.ValidateAll() if the designated constraints
// aren't met.
type AccountBalanceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AccountBalanceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AccountBalanceMultiError) AllErrors() []error { return m }

// AccountBalanceValidationError is the validation error returned by
// AccountBalance.Validate if the designated constraints aren't met.
type AccountBalanceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AccountBalanceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AccountBalanceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AccountBalanceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AccountBalanceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AccountBalanceValidationError) ErrorName() string { return "AccountBalanceValidationError" }

// Error satisfies the builtin error interface
func (e AccountBalanceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAccountBalance.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AccountBalanceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AccountBalanceValidationError{}

// Validate checks the field values on AccountAndBalance with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AccountAndBalance) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AccountAndBalance with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AccountAndBalanceMultiError, or nil if none found.
func (m *AccountAndBalance) ValidateAll() error {
	return m.validate(true)
}

func (m *AccountAndBalance) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AccountId

	// no validation rules for Currency

	// no validation rules for Balance

	// no validation rules for Hold

	// no validation rules for Available

	if all {
		switch v := interface{}(m.GetBalanceAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AccountAndBalanceValidationError{
					field:  "BalanceAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AccountAndBalanceValidationError{
					field:  "BalanceAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBalanceAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccountAndBalanceValidationError{
				field:  "BalanceAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return AccountAndBalanceMultiError(errors)
	}

	return nil
}

// AccountAndBalanceMultiError is an error wrapping multiple validation errors
// returned by AccountAndBalance.ValidateAll() if the designated constraints
// aren't met.
type AccountAndBalanceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AccountAndBalanceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AccountAndBalanceMultiError) AllErrors() []error { return m }

// AccountAndBalanceValidationError is the validation error returned by
// AccountAndBalance.Validate if the designated constraints aren't met.
type AccountAndBalanceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AccountAndBalanceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AccountAndBalanceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AccountAndBalanceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AccountAndBalanceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AccountAndBalanceValidationError) ErrorName() string {
	return "AccountAndBalanceValidationError"
}

// Error satisfies the builtin error interface
func (e AccountAndBalanceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAccountAndBalance.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AccountAndBalanceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AccountAndBalanceValidationError{}

// Validate checks the field values on Transaction with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Transaction) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Transaction with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TransactionMultiError, or
// nil if none found.
func (m *Transaction) ValidateAll() error {
	return m.validate(true)
}

func (m *Transaction) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for SenderId

	// no validation rules for ReceiverId

	// no validation rules for RequestId

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TransactionValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TransactionValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TransactionValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for TransactionType

	// no validation rules for TransactionStatus

	if m.FinalizedAt != nil {

		if all {
			switch v := interface{}(m.GetFinalizedAt()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TransactionValidationError{
						field:  "FinalizedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TransactionValidationError{
						field:  "FinalizedAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetFinalizedAt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TransactionValidationError{
					field:  "FinalizedAt",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return TransactionMultiError(errors)
	}

	return nil
}

// TransactionMultiError is an error wrapping multiple validation errors
// returned by Transaction.ValidateAll() if the designated constraints aren't met.
type TransactionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TransactionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TransactionMultiError) AllErrors() []error { return m }

// TransactionValidationError is the validation error returned by
// Transaction.Validate if the designated constraints aren't met.
type TransactionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TransactionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TransactionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TransactionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TransactionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TransactionValidationError) ErrorName() string { return "TransactionValidationError" }

// Error satisfies the builtin error interface
func (e TransactionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTransaction.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TransactionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TransactionValidationError{}

// Validate checks the field values on Hold with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Hold) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Hold with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in HoldMultiError, or nil if none found.
func (m *Hold) ValidateAll() error {
	return m.validate(true)
}

func (m *Hold) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for AccountId

	// no validation rules for TransactionId

	// no validation rules for Amount

	// no validation rules for Direction

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HoldValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HoldValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HoldValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetReleasedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HoldValidationError{
					field:  "ReleasedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HoldValidationError{
					field:  "ReleasedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetReleasedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HoldValidationError{
				field:  "ReleasedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return HoldMultiError(errors)
	}

	return nil
}

// HoldMultiError is an error wrapping multiple validation errors returned by
// Hold.ValidateAll() if the designated constraints aren't met.
type HoldMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HoldMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HoldMultiError) AllErrors() []error { return m }

// HoldValidationError is the validation error returned by Hold.Validate if the
// designated constraints aren't met.
type HoldValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HoldValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HoldValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HoldValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HoldValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HoldValidationError) ErrorName() string { return "HoldValidationError" }

// Error satisfies the builtin error interface
func (e HoldValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHold.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HoldValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HoldValidationError{}

// Validate checks the field values on Entry with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Entry) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Entry with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in EntryMultiError, or nil if none found.
func (m *Entry) ValidateAll() error {
	return m.validate(true)
}

func (m *Entry) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for AccountId

	// no validation rules for TransactionId

	// no validation rules for Amount

	// no validation rules for Direction

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EntryValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EntryValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EntryValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return EntryMultiError(errors)
	}

	return nil
}

// EntryMultiError is an error wrapping multiple validation errors returned by
// Entry.ValidateAll() if the designated constraints aren't met.
type EntryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EntryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EntryMultiError) AllErrors() []error { return m }

// EntryValidationError is the validation error returned by Entry.Validate if
// the designated constraints aren't met.
type EntryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EntryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EntryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EntryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EntryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EntryValidationError) ErrorName() string { return "EntryValidationError" }

// Error satisfies the builtin error interface
func (e EntryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEntry.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EntryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EntryValidationError{}