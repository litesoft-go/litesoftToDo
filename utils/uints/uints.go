package uints

import (
	"math"
	"math/bits"
	"errors"
)

const SIZE = bits.UintSize // 32 or 64

func IsBetweenInclusive(pMin, pToCheck, pMax uint) bool {
	return (pMin <= pToCheck) && (pToCheck <= pMax)
}

func IsEven(pToCheck uint) bool {
	return (pToCheck & 1) == 0
}

func Min(pValue1, pValue2 uint) uint {
	if pValue1 < pValue2 {
		return pValue1
	}
	return pValue2
}

func Max(pValue1, pValue2 uint) uint {
	if pValue1 > pValue2 {
		return pValue1
	}
	return pValue2
}

func FromInt64(pValue int64) (rValue uint, err error) {
	if pValue < 0 {
		err = errors.New("int64 value Negative")
		return
	}
	if SIZE == 32 {
		if pValue > math.MaxUint32 {
			err = errors.New("int64 value Too Large (> 31 bits)")
			return
		}
	}
	rValue = uint(pValue)
	return
}

func FromUint64(pValue uint64) (rValue uint, err error) {
	if SIZE == 32 {
		if pValue > math.MaxUint32 {
			err = errors.New("uint64 value Too Large (> 32 bits)")
			return
		}
	}
	rValue = uint(pValue)
	return
}

