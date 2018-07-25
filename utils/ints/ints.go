package ints

import (
	"math"
	"math/bits"
	"errors"
)

const SIZE = bits.UintSize // 32 or 64

func IsBetweenInclusive(pMin, pToCheck, pMax int) bool {
	return (pMin <= pToCheck) && (pToCheck <= pMax)
}

func IsEven(pToCheck int) bool {
	return (pToCheck & 1) == 0
}

func Min(pValue1, pValue2 int) int {
	if pValue1 < pValue2 {
		return pValue1
	}
	return pValue2
}

func Max(pValue1, pValue2 int) int {
	if pValue1 > pValue2 {
		return pValue1
	}
	return pValue2
}

func FromInt64(pValue int64) (rValue int, err error) {
	if SIZE == 32 {
		if pValue > math.MaxInt32 {
			err = errors.New("int64 value Too Large (> 31 bits)")
			return
		}
		if pValue < math.MinInt32 {
			err = errors.New("int64 value Too Small (negative)")
			return
		}
	}
	rValue = int(pValue)
	return
}

func FromUint64(pValue uint64) (rValue int, err error) {
	if SIZE == 32 {
		if pValue > math.MaxInt32 {
			err = errors.New("uint64 value Too Large (> 31 bits)")
			return
		}
	} else if pValue > math.MaxInt64 {
		err = errors.New("uint64 value Too Large (> 63 bits)")
		return
	}
	rValue = int(pValue)
	return
}

