// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: file/service/v1/attachment.proto

package servicev1

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

// Validate checks the field values on Attachment with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Attachment) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Attachment with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AttachmentMultiError, or
// nil if none found.
func (m *Attachment) ValidateAll() error {
	return m.validate(true)
}

func (m *Attachment) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if m.Name != nil {
		// no validation rules for Name
	}

	if m.Path != nil {
		// no validation rules for Path
	}

	if m.FileKey != nil {
		// no validation rules for FileKey
	}

	if m.ThumbPath != nil {
		// no validation rules for ThumbPath
	}

	if m.MediaType != nil {
		// no validation rules for MediaType
	}

	if m.Suffix != nil {
		// no validation rules for Suffix
	}

	if m.Width != nil {
		// no validation rules for Width
	}

	if m.Height != nil {
		// no validation rules for Height
	}

	if m.Size != nil {
		// no validation rules for Size
	}

	if m.Type != nil {
		// no validation rules for Type
	}

	if m.CreateTime != nil {

		if all {
			switch v := interface{}(m.GetCreateTime()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AttachmentValidationError{
						field:  "CreateTime",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AttachmentValidationError{
						field:  "CreateTime",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetCreateTime()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AttachmentValidationError{
					field:  "CreateTime",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.UpdateTime != nil {

		if all {
			switch v := interface{}(m.GetUpdateTime()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AttachmentValidationError{
						field:  "UpdateTime",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AttachmentValidationError{
						field:  "UpdateTime",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetUpdateTime()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AttachmentValidationError{
					field:  "UpdateTime",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.DeleteTime != nil {

		if all {
			switch v := interface{}(m.GetDeleteTime()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AttachmentValidationError{
						field:  "DeleteTime",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AttachmentValidationError{
						field:  "DeleteTime",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetDeleteTime()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AttachmentValidationError{
					field:  "DeleteTime",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return AttachmentMultiError(errors)
	}

	return nil
}

// AttachmentMultiError is an error wrapping multiple validation errors
// returned by Attachment.ValidateAll() if the designated constraints aren't met.
type AttachmentMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AttachmentMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AttachmentMultiError) AllErrors() []error { return m }

// AttachmentValidationError is the validation error returned by
// Attachment.Validate if the designated constraints aren't met.
type AttachmentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AttachmentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AttachmentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AttachmentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AttachmentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AttachmentValidationError) ErrorName() string { return "AttachmentValidationError" }

// Error satisfies the builtin error interface
func (e AttachmentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttachment.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AttachmentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AttachmentValidationError{}

// Validate checks the field values on ListAttachmentResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListAttachmentResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListAttachmentResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListAttachmentResponseMultiError, or nil if none found.
func (m *ListAttachmentResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListAttachmentResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListAttachmentResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListAttachmentResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListAttachmentResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return ListAttachmentResponseMultiError(errors)
	}

	return nil
}

// ListAttachmentResponseMultiError is an error wrapping multiple validation
// errors returned by ListAttachmentResponse.ValidateAll() if the designated
// constraints aren't met.
type ListAttachmentResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListAttachmentResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListAttachmentResponseMultiError) AllErrors() []error { return m }

// ListAttachmentResponseValidationError is the validation error returned by
// ListAttachmentResponse.Validate if the designated constraints aren't met.
type ListAttachmentResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListAttachmentResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListAttachmentResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListAttachmentResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListAttachmentResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListAttachmentResponseValidationError) ErrorName() string {
	return "ListAttachmentResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListAttachmentResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListAttachmentResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListAttachmentResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListAttachmentResponseValidationError{}

// Validate checks the field values on GetAttachmentRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetAttachmentRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetAttachmentRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetAttachmentRequestMultiError, or nil if none found.
func (m *GetAttachmentRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetAttachmentRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetAttachmentRequestMultiError(errors)
	}

	return nil
}

// GetAttachmentRequestMultiError is an error wrapping multiple validation
// errors returned by GetAttachmentRequest.ValidateAll() if the designated
// constraints aren't met.
type GetAttachmentRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetAttachmentRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetAttachmentRequestMultiError) AllErrors() []error { return m }

// GetAttachmentRequestValidationError is the validation error returned by
// GetAttachmentRequest.Validate if the designated constraints aren't met.
type GetAttachmentRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAttachmentRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAttachmentRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAttachmentRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAttachmentRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAttachmentRequestValidationError) ErrorName() string {
	return "GetAttachmentRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetAttachmentRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAttachmentRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAttachmentRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAttachmentRequestValidationError{}

// Validate checks the field values on CreateAttachmentRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateAttachmentRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateAttachmentRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateAttachmentRequestMultiError, or nil if none found.
func (m *CreateAttachmentRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateAttachmentRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetAttachment()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateAttachmentRequestValidationError{
					field:  "Attachment",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateAttachmentRequestValidationError{
					field:  "Attachment",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAttachment()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateAttachmentRequestValidationError{
				field:  "Attachment",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.OperatorId != nil {
		// no validation rules for OperatorId
	}

	if len(errors) > 0 {
		return CreateAttachmentRequestMultiError(errors)
	}

	return nil
}

// CreateAttachmentRequestMultiError is an error wrapping multiple validation
// errors returned by CreateAttachmentRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateAttachmentRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateAttachmentRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateAttachmentRequestMultiError) AllErrors() []error { return m }

// CreateAttachmentRequestValidationError is the validation error returned by
// CreateAttachmentRequest.Validate if the designated constraints aren't met.
type CreateAttachmentRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateAttachmentRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateAttachmentRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateAttachmentRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateAttachmentRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateAttachmentRequestValidationError) ErrorName() string {
	return "CreateAttachmentRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateAttachmentRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateAttachmentRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateAttachmentRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateAttachmentRequestValidationError{}

// Validate checks the field values on UpdateAttachmentRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateAttachmentRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateAttachmentRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateAttachmentRequestMultiError, or nil if none found.
func (m *UpdateAttachmentRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateAttachmentRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetAttachment()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateAttachmentRequestValidationError{
					field:  "Attachment",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateAttachmentRequestValidationError{
					field:  "Attachment",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAttachment()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateAttachmentRequestValidationError{
				field:  "Attachment",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.OperatorId != nil {
		// no validation rules for OperatorId
	}

	if len(errors) > 0 {
		return UpdateAttachmentRequestMultiError(errors)
	}

	return nil
}

// UpdateAttachmentRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateAttachmentRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateAttachmentRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateAttachmentRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateAttachmentRequestMultiError) AllErrors() []error { return m }

// UpdateAttachmentRequestValidationError is the validation error returned by
// UpdateAttachmentRequest.Validate if the designated constraints aren't met.
type UpdateAttachmentRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateAttachmentRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateAttachmentRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateAttachmentRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateAttachmentRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateAttachmentRequestValidationError) ErrorName() string {
	return "UpdateAttachmentRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateAttachmentRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateAttachmentRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateAttachmentRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateAttachmentRequestValidationError{}

// Validate checks the field values on DeleteAttachmentRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteAttachmentRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteAttachmentRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteAttachmentRequestMultiError, or nil if none found.
func (m *DeleteAttachmentRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteAttachmentRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if m.OperatorId != nil {
		// no validation rules for OperatorId
	}

	if len(errors) > 0 {
		return DeleteAttachmentRequestMultiError(errors)
	}

	return nil
}

// DeleteAttachmentRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteAttachmentRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteAttachmentRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteAttachmentRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteAttachmentRequestMultiError) AllErrors() []error { return m }

// DeleteAttachmentRequestValidationError is the validation error returned by
// DeleteAttachmentRequest.Validate if the designated constraints aren't met.
type DeleteAttachmentRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteAttachmentRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteAttachmentRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteAttachmentRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteAttachmentRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteAttachmentRequestValidationError) ErrorName() string {
	return "DeleteAttachmentRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteAttachmentRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteAttachmentRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteAttachmentRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteAttachmentRequestValidationError{}
