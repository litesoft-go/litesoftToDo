package validate

import (
	"litesoftToDo/utils/validate/stringCheck"
	"strings"
)

type SourcedString struct {
	CommonSourced
	mValue string
}

func NewSourcedString(pSourcedData SourcedData) (rInstance *SourcedString) {
	rInstance = &SourcedString{CommonSourced: CommonSourced{mSourcedData:pSourcedData}}
	if rInstance.Found() {
		rInstance.mValue = rInstance.getRawString()
	}
	return
}

func (this *SourcedString) TrimSpace() *SourcedString {
	if this != nil {
		this.mValue = strings.TrimSpace(this.mValue)
	}
	return this
}

func (this *SourcedString) Required() *SourcedString {
	return this.check(stringCheck.Required)
}

func (this *SourcedString) NoWhiteSpace() *SourcedString {
	return this.check(stringCheck.NoWhiteSpace)
}

func (this *SourcedString) NoLeadingOrTrailingNorAdjacentWhiteSpace() *SourcedString {
	return this.check(stringCheck.NoLeadingOrTrailingNorAdjacentWhiteSpace)
}

func (this *SourcedString) check(pFunc func(string) error) *SourcedString {
	if !this.HasError() {
		this.SetError(pFunc(this.mValue))
	}
	return this
}