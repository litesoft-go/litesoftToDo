package validation

import (
	"fmt"
	"litesoftToDo/utils/berrors"
)

func IsValidation(err error) bool {
	type validation interface {
		IsValidationError() bool
	}
	for ; err != nil; err = berrors.GetCause(err) {
		verr, ok := err.(validation)
		if ok && verr.IsValidationError() {
			return true
		}
	}
	return false
}

type Error struct {
	mMsg   string
	mCause error
}

func NewError(pMsg string) *Error {
	return &Error{mMsg: pMsg}
}

func NewErrorf(pFormat string, args ...interface{}) *Error {
	return NewError(fmt.Sprintf(pFormat, args...))
}

func (this *Error) WithCause(err error) error {
	if this != nil {
		this.mCause = err
	}
	return this
}

func (this *Error) Cause() (err error) {
	if this != nil {
		err = this.mCause
	}
	return
}

func (this *Error) Error() (rMsg string) {
	if this != nil {
		rMsg = this.mMsg
		if this.mCause != nil {
			rMsg += ": " + this.Cause().Error()
		}
	}
	return
}

func (this *Error) IsValidationError() bool {
	return true
}
