package validate

import (
	"strings"
	"unicode"
	"litesoftToDo/utils/validation"
)

func RequiredNoWhiteSpace(pWhat, pToCheck string) error {
	if len(pToCheck) == 0 {
		return required(pWhat)
	}
	return NoWhiteSpace(pWhat, pToCheck)
}

func NoWhiteSpace(pWhat, pToCheck string) (err error) {
	if -1 != strings.IndexFunc(pToCheck, unicode.IsSpace) {
		err = validation.NewErrorf("%s has whitespace: '%s'", pWhat, pToCheck)
	}
	return
}

func RequiredAndNoLeadingNorMultipleInnerWhiteSpace(pWhat, pToCheck string) error {
	if len(pToCheck) == 0 {
		return required(pWhat)
	}
	return NoLeadingNorMultipleInnerWhiteSpace(pWhat, pToCheck)
}

func NoLeadingNorMultipleInnerWhiteSpace(pWhat, pToCheck string) (err error) {
	if len(pToCheck) != 0 {
		zRunes := []rune(pToCheck) // NOT Empty - see above
		if unicode.IsSpace(zRunes[0]) || unicode.IsSpace(zRunes[len(zRunes)-1]) {
			return validation.NewErrorf("%s has leading or trailing whitespace: '%s'", pWhat, pToCheck)
		}
		zLastWhitespaceAt := -1
		for i, zRune := range zRunes {
			if unicode.IsSpace(zRune) {
				if (zLastWhitespaceAt + 1) == i { // adjacent
					return validation.NewErrorf("%s has adjacent (at %d & %d) whitespace: '%s'",
						pWhat, zLastWhitespaceAt, i, pToCheck)
				}
				zLastWhitespaceAt = i
			}
		}
	}
	return
}

func required(pWhat string) error {
	return validation.NewError(pWhat + " is REQUIRED, but was empty")
}
