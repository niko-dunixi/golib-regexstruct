package regexstruct

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseMatchMap struct {
	regex    regexp.Regexp
	input    string
	expected map[string]string
}

var (
	matchMapTestCases = []TestCaseMatchMap{
		{
			regex: *regexp.MustCompile(`/(?P<first>\w+)/(?P<second>\w+)`),
			input: "/this_is_first/this_is_second",
			expected: map[string]string{
				"first":  "this_is_first",
				"second": "this_is_second",
			},
		},
	}
)

func TestRegexMatchMap(t *testing.T) {
	for _, tc := range matchMapTestCases {
		tc := tc
		name := fmt.Sprintf("%s %s", tc.regex.String(), tc.input)
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			result, isMatch := RegexMatchMap(&tc.regex, tc.input)
			assert.True(t, isMatch)
			if assert.Len(t, result, len(tc.expected)) {
				for expectedKey, expectedValue := range tc.expected {
					if assert.Contains(t, result, expectedKey) {
						assert.Equal(t, result[expectedKey], expectedValue)
					}
				}
			}
		})
	}
}
