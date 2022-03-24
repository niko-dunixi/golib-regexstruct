//go:build go1.18
// +build go1.18

package regexstruct

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegexMatch(t *testing.T) {
	regex := regexp.MustCompile(`/(?P<first>\w+)/(?P<second>\w+)`)
	input := "/this_is_first/this_is_second"
	expected := struct {
		First  string
		Second string
	}{
		First:  "this_is_first",
		Second: "this_is_second",
	}
	actual := struct {
		First  string
		Second string
	}{}
	err := RegexMatch(regex, input, &actual)
	assert.NoError(t, err)

	assert.Equal(t, expected.First, actual.First)
	assert.Equal(t, expected.Second, actual.Second)
}
