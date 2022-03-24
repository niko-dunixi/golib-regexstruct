//go:build go1.18
// +build go1.18

package regexstruct

import (
	"fmt"
	"regexp"

	"github.com/mitchellh/mapstructure"
)

type NoMatchError struct {
	regex *regexp.Regexp
	input string
}

func (nme NoMatchError) Error() string {
	pattern := nme.regex.String()
	return fmt.Sprintf("the regex (%s) did not match the input string (%s)", pattern, nme.input)
}

func RegexMatch[T any](r *regexp.Regexp, s string, item *T) error {
	matchedMap, isMatch := RegexMatchMap(r, s)
	if !isMatch {
		return NoMatchError{
			regex: r,
			input: s,
		}
	}
	return mapstructure.Decode(matchedMap, item)
}
