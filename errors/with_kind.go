package errors

import (
	"fmt"

	"github.com/pkg/errors"
)

type Kind string

const (
	KindCommunication  Kind = "communication-error"
	KindDataAccess     Kind = "data-access-error"
	KindDataProcessing Kind = "data-processing-error"
)

type ErrorWithKind interface {
	error
	fmt.Stringer
	Cause() error

	Kind() Kind
}

type withKindError struct {
	kind    Kind
	message string
	cause   error
}

func WithKindError(kind Kind, message string, parent error) ErrorWithKind {
	return &withKindError{
		kind:    kind,
		message: message,
		cause:   parent,
	}
}

func Communication(err error, message string) ErrorWithKind {
	return WithKindError(KindCommunication, message, err)
}

func DataAccess(err error, message string) ErrorWithKind {
	return WithKindError(KindDataAccess, message, err)
}

func DataProcessing(err error, message string) ErrorWithKind {
	return WithKindError(KindDataProcessing, message, err)
}

func (err *withKindError) Error() string {
	return err.String()
}

func (err *withKindError) String() string {
	return fmt.Sprintf("%s<%s: %s>", err.kind, err.message, err.cause)
}

func (err *withKindError) Cause() error {
	return err.cause
}

func (err *withKindError) Kind() Kind {
	return err.kind
}

func Cause(err error) error {
	return errors.Cause(err)
}
