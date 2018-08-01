package validate

import (
	"strings"
	"fmt"
	"litesoftToDo/utils/stringer"
)

type SourcedInfo struct {
	mSource, mReference string
}

func NewSourcedInfo(pSource, pReference string) *SourcedInfo {
	info := &SourcedInfo{}
	info.Populate(pSource, pReference)
	return info
}

func (this *SourcedInfo) Populate(pSource, pReference string) {
	this.mSource = strings.TrimSpace(pSource)
	this.mReference = requireSignificant(pReference)
}

func (this *SourcedInfo) String() (rInfo string) {
	rInfo = "'" + this.mReference + "'"
	if len(this.mSource) != 0 {
		rInfo = this.mSource + " " + rInfo
	}
	return
}

func (this *SourcedInfo) WithData(pRawValue string, pFound bool) *BasicSourcedData {
	return &BasicSourcedData{mSource: this, mRawValue: pRawValue, mFound: pFound}
}

type BasicSourcedData struct {
	mSource   fmt.Stringer
	mRawValue string
	mFound    bool
}

func (this *BasicSourcedData) Found() bool {
	return (this != nil) && this.mFound // Left to Right
}

func (this *BasicSourcedData) Get() (rRawValue string) {
	if this != nil {
		rRawValue = this.mRawValue
	}
	return
}

func (this *BasicSourcedData) Source() fmt.Stringer {
	if this == nil {
		return stringer.Empty
	}
	return this.mSource
}

func (this *BasicSourcedData) String() string {
	return stringOfSourcedData(this)
}

func (this *BasicSourcedData) Named(pLabel string) *NamedSourceData {
	return &NamedSourceData{mSource: stringer.Augment(this.Source(), ": ", pLabel), mRawValue: this.Get(), mFound: this.Found()}
}

type NamedSourceData struct {
	mSource   fmt.Stringer
	mRawValue string
	mFound    bool
}

func (this *NamedSourceData) Found() bool {
	return (this != nil) && this.mFound // Left to Right
}

func (this *NamedSourceData) Get() (rRawValue string) {
	if this != nil {
		rRawValue = this.mRawValue
	}
	return
}

func (this *NamedSourceData) Source() fmt.Stringer {
	if this == nil {
		return stringer.Empty
	}
	return this.mSource
}

func (this *NamedSourceData) String() string {
	return stringOfSourcedData(this)
}

func (this *NamedSourceData) AndName(pLabel string) *NamedSourceData {
	return &NamedSourceData{mSource: stringer.Augment(this.Source(), "/", pLabel), mRawValue: this.Get(), mFound: this.Found()}
}

func requireSignificant(pWhat string) (rWhat string) {
	rWhat = strings.TrimSpace(pWhat)
	err := RequiredNoWhiteSpace("What", rWhat)
	if err != nil {
		panic(err)
	}
	return
}

func stringOfSourcedData(this SourcedData) (rResults string) {
	if this.Found() {
		rResults = "'" + this.Get() + "'"
	} else {
		rResults = "(Not Found)"
	}
	zSource := this.Source().String()
	if len(zSource) != 0 {
		rResults = zSource + ": " + rResults
	}
	return
}