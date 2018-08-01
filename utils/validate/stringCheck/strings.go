package stringCheck

import (
	"strings"
	"unicode"
	"litesoftToDo/utils/validation"
)

func Required(pToCheck string) (err error) {
	if len(pToCheck) == 0 {
		err = validation.NewError(" is REQUIRED, but was empty")
	}
	return
}

func NoWhiteSpace(pToCheck string) (err error) {
	if -1 != strings.IndexFunc(pToCheck, unicode.IsSpace) {
		err = validation.NewErrorf("has whitespace: '%s'", pToCheck)
	}
	return
}

func NoLeadingOrTrailingNorAdjacentWhiteSpace(pToCheck string) (err error) {
	if len(pToCheck) != 0 {
		zRunes := []rune(pToCheck) // NOT Empty - see above
		if unicode.IsSpace(zRunes[0]) || unicode.IsSpace(zRunes[len(zRunes)-1]) {
			return validation.NewErrorf("has leading or trailing whitespace: '%s'", pToCheck)
		}
		zLastWhitespaceAt := -1
		for i, zRune := range zRunes {
			if unicode.IsSpace(zRune) {
				if (zLastWhitespaceAt + 1) == i { // adjacent
					return validation.NewErrorf("has adjacent (at %d & %d) whitespace: '%s'",
						zLastWhitespaceAt, i, pToCheck)
				}
				zLastWhitespaceAt = i
			}
		}
	}
	return
}
