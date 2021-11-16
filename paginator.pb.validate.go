// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: paginator.proto

package paginator

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

// Validate checks the field values on Paginator with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Paginator) Validate() error {
	if m == nil {
		return nil
	}

	if val := m.GetPage(); val < 1 || val > 1000 {
		return PaginatorValidationError{
			field:  "Page",
			reason: "value must be inside range [1, 1000]",
		}
	}

	if val := m.GetLimit(); val < 1 || val > 100 {
		return PaginatorValidationError{
			field:  "Limit",
			reason: "value must be inside range [1, 100]",
		}
	}

	// no validation rules for Total

	return nil
}

// PaginatorValidationError is the validation error returned by
// Paginator.Validate if the designated constraints aren't met.
type PaginatorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PaginatorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PaginatorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PaginatorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PaginatorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PaginatorValidationError) ErrorName() string { return "PaginatorValidationError" }

// Error satisfies the builtin error interface
func (e PaginatorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPaginator.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PaginatorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PaginatorValidationError{}
