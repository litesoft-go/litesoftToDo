package uuid

import (
	"litesoftToDo/utils/ints"
	"fmt"
	"strconv"
)

/*
 * Implementations (that are not mocked) are expected to generate the RFC 4122 36 byte String form with 4 dashes
 */
type UUIDsourceMock struct {
	mPrefix      string
	mDigitsToAdd int
	mNext        int
}

func NewMock(pPrefix string, pDesiredLength int) UUIDsourceFactory {
	if !ints.IsBetweenInclusive(1, pDesiredLength, 36) {
		panic(fmt.Sprintf("DesiredLength not: 1 <= %d <= 36", pDesiredLength))
	}
	zDigitsToAdd := pDesiredLength - len(pPrefix)
	if zDigitsToAdd < 1 {
		panic(fmt.Sprintf("DesiredLength (%d) not longer than Prefix: '%s'", pDesiredLength, pPrefix))
	}
	return &UUIDsourceMock{mPrefix: pPrefix, mDigitsToAdd: zDigitsToAdd}
}

func (this *UUIDsourceMock) GetSource() UUIDsource {
	return this.asSource
}

func (this *UUIDsourceMock) asSource() (rUUID string, err error) {
	zNext := strconv.Itoa(this.mNext)
	if len(zNext) > this.mDigitsToAdd {
		err = fmt.Errorf("next Suffix ('%s') too long to generate the DesiredLength", zNext)
		return
	}
	for len(zNext) < this.mDigitsToAdd {
		zNext = "0" + zNext
	}
	rUUID = this.mPrefix + zNext
	this.mNext++
	return
}
