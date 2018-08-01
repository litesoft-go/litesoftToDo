package stringer

import "fmt"

type instance struct{}

var Empty instance

func (this instance) String() string {
	return ""
}

type chain struct {
	mPrevious         fmt.Stringer
	mSeperator, mPlus string
}

func Augment(pExisting fmt.Stringer, pSeperator, pPlus string) fmt.Stringer {
	if pExisting == nil {
		pExisting = Empty
		pSeperator = ""
	}
	return &chain{mPrevious:pExisting, mSeperator:pSeperator, mPlus:pPlus}
}

func (this *chain) String() string {
	if this == nil {
		return ""
	}
	return this.mPrevious.String() + this.mSeperator + this.mPlus
}
