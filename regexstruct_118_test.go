//go:build go1.18
// +build go1.18

package regexstruct

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

type example struct {
	first  string
	second string
}

func (e *example) Unmarshal(m map[string]string) error {
	e.first = m["first"]
	e.second = m["second"]
	return nil
}

func TestRegexMatch(t *testing.T) {
	regex := regexp.MustCompile(`/(?P<first>\w+)/(?P<second>\w+)`)
	input := "/this_is_first/this_is_second"
	expected := example{
		first:  "this_is_first",
		second: "this_is_second",
	}
	actual := example{}
	err := RegexMatch(regex, input, &actual)
	assert.NoError(t, err)

	assert.Equal(t, expected.first, actual.first)
	assert.Equal(t, expected.second, actual.second)
}
