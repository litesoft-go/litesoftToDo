package validate

import (
	"strings"
	"unicode"
	"fmt"
	"errors"
)

func NoWhiteSpace(pWhat, pToCheck string) (err error) {
	if -1 != strings.IndexFunc(pToCheck, unicode.IsSpace) {
		err = fmt.Errorf("%s has whitespace: '%s'", pWhat, pToCheck)
	}
	return
}

func RequiredAndNoLeadingNorMultipleInnerWhiteSpace(pWhat, pToCheck string) error {
	if len(pToCheck) == 0 {
		return errors.New(pWhat + " is REQUIRED, but was empty")
	}
	return NoLeadingNorMultipleInnerWhiteSpace(pWhat, pToCheck)
}

func NoLeadingNorMultipleInnerWhiteSpace(pWhat, pToCheck string) (err error) {
	if len(pToCheck) != 0 {
		zRunes := []rune(pToCheck) // NOT Empty - see above
		if unicode.IsSpace(zRunes[0]) || unicode.IsSpace(zRunes[len(zRunes)-1]) {
			return fmt.Errorf("%s has leading or trailing whitespace: '%s'", pWhat, pToCheck)
		}
		zLastWhitespaceAt := -1
		for i, zRune := range zRunes {
			if unicode.IsSpace(zRune) {
				if (zLastWhitespaceAt + 1) == i { // adjacent
					return fmt.Errorf("%s has adjacent (at %d & %d) whitespace: '%s'",
						zLastWhitespaceAt, i, pWhat, pToCheck)
				}
				zLastWhitespaceAt = i
			}
		}
	}
	return
	}
