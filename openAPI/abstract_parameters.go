package openAPI

import (
	"strings"
	"strconv"
	"github.com/pkg/errors"
	"litesoftToDo/utils/ints"
	"litesoftToDo/utils/uints"
)

type ParameterSource func(pKey string) (string)

type AbstractParameters struct {
	ParamSource ParameterSource
}

func (this *AbstractParameters) getRaw(pKey string) (rFound string) {
	if this != nil {
		rFound = strings.TrimSpace(this.ParamSource(pKey))
	}
	return
}

func (this *AbstractParameters) getRawNonEmpty(pKey string) (rFound string, err error) {
	rFound = this.getRaw(pKey)
	if len(rFound) == 0 {
		err = errors.New("No Value or value empty for key: " + pKey)
	}
	return
}

func (this *AbstractParameters) GetString(pKey string) string {
	return this.getRaw(pKey)
}

func (this *AbstractParameters) GetStringOr(pKey, pOr string) (rFound string) {
	rFound = this.getRaw(pKey)
	if len(rFound) == 0 {
		rFound = pOr
	}
	return
}

func (this *AbstractParameters) GetInt(pKey string) (rValue int, err error) {
	zValue, err := this.GetInt64(pKey)
	if err == nil {
		rValue, err = ints.FromInt64(zValue)
	}
	return
}

func (this *AbstractParameters) GetIntOr(pKey string, pOr int) (rValue int) {
	zValue, err := this.GetInt(pKey)
	if err == nil {
		rValue = zValue
	} else {
		rValue = pOr
	}
	return
}

func (this *AbstractParameters) GetInt64(pKey string) (rValue int64, err error) {
	zString, err := this.getRawNonEmpty(pKey) // Note: mix existence w/ :=
	if err == nil {
		rValue, err = strconv.ParseInt(zString, 10, 64)
	}
	return
}

func (this *AbstractParameters) GetInt64Or(pKey string, pOr int64) (rValue int64) {
	zValue, err := this.GetInt64(pKey)
	if err == nil {
		rValue = zValue
	}
	return
}

func (this *AbstractParameters) GetUint(pKey string) (rValue uint, err error) {
	zValue, err := this.GetUint64(pKey)
	if err == nil {
		rValue, err = uints.FromUint64(zValue)
	}
	return
}

func (this *AbstractParameters) GetUintOr(pKey string, pOr uint) (rValue uint) {
	zValue, err := this.GetUint(pKey)
	if err == nil {
		rValue = zValue
	} else {
		rValue = pOr
	}
	return
}

func (this *AbstractParameters) GetUint64(pKey string) (rValue uint64, err error) {
	zString, err := this.getRawNonEmpty(pKey) // Note: mix existence w/ :=
	if err == nil {
		rValue, err = strconv.ParseUint(zString, 10, 64)
	}
	return
}

func (this *AbstractParameters) GetUint64Or(pKey string, pOr uint64) (rValue uint64) {
	zValue, err := this.GetUint64(pKey)
	if err == nil {
		rValue = zValue
	}
	return
}
