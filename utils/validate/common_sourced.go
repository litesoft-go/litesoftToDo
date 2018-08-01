package validate

import (
	"fmt"
	"litesoftToDo/utils/berrors"
)

type CommonSourced struct {
	mSourcedData SourcedData
	mErr         error
}

func (this *CommonSourced) GetSourceData() (rSourcedData SourcedData) {
	if this != nil {
		rSourcedData = this.mSourcedData
	}
	return
}

func (this *CommonSourced) HasError() bool {
	return (this != nil) && (this.mErr != nil) // Left to Right!
}

func (this *CommonSourced) GetError() (err error) {
	if this != nil {
		err = this.mErr
	}
	return
}

func (this *CommonSourced) SetError(err error) {
	if this != nil {
		this.mErr = this.augmentError(err)
	}
	return
}

func (this *CommonSourced) Found() (rFound bool) {
	zSourcedData := this.GetSourceData()
	if zSourcedData != nil {
		rFound = zSourcedData.Found()
	}
	return
}

func (this *CommonSourced) Source() (rSource fmt.Stringer) {
	zSourcedData := this.GetSourceData()
	if zSourcedData != nil {
		rSource = zSourcedData.Source()
	}
	return
}

func (this *CommonSourced) getRawString() (rRawString string) {
	zSourcedData := this.GetSourceData()
	if zSourcedData != nil {
		rRawString = zSourcedData.Get()
	}
	return
}

func (this *CommonSourced) augmentError(pError error) (err error) {
	if (this == nil) || pError == nil {
		return
	}
	zSource := getSourced(pError)
	if zSource == this.Source() {
		err = pError
	} else {
		err = newSourcedError(this.Source(), pError)
	}
	return
}

func getSourced(err error) (rSource fmt.Stringer) {
	type sourced interface {
		GetSource() fmt.Stringer
	}
	for ; err != nil; err = berrors.GetCause(err) {
		serr, ok := err.(sourced)
		if ok {
			return serr.GetSource()
		}
	}
	return
}

type sourcedError struct {
	mSource fmt.Stringer
	mCause  error
}

func newSourcedError(pSource fmt.Stringer, pCause error) *sourcedError {
	return &sourcedError{mSource: pSource, mCause:pCause}
}

func (this *sourcedError) Cause() (err error) {
	if this != nil {
		err = this.mCause
	}
	return
}

func (this *sourcedError) Error() (rMsg string) {
	if this != nil {
		rMsg = this.mSource.String() + ": " + this.Cause().Error()
	}
	return
}

func (this *sourcedError) GetSource() fmt.Stringer {
	return this.mSource
}
