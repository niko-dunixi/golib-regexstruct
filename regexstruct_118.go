//go:build go1.18
// +build go1.18

package regexstruct

import (
	"fmt"
	"regexp"
)

type NoMatchError struct {
	regex *regexp.Regexp
	input string
}

func (nme NoMatchError) Error() string {
	pattern := nme.regex.String()
	return fmt.Sprintf("the regex (%s) did not match the input string (%s)", pattern, nme.input)
}

type Unmarshaler interface {
	Unmarshal(map[string]string) error
}

func RegexMatch[T Unmarshaler](r *regexp.Regexp, s string, item T) error {
	matchedMap, isMatch := RegexMatchMap(r, s)
	if !isMatch {
		return NoMatchError{
			regex: r,
			input: s,
		}
	}
	return item.Unmarshal(matchedMap)
}
