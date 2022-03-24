package regexstruct

import (
	"regexp"
)

func RegexMatchMap(r *regexp.Regexp, s string) (map[string]string, bool) {
	groupNames := r.SubexpNames()
	matches := r.FindStringSubmatch(s)
	if nameCount := len(groupNames); nameCount != len(matches) || nameCount == 0 {
		return nil, false
	}
	groupNames = groupNames[1:]
	matches = matches[1:]
	captureMap := make(map[string]string)
	for index, currentGroupName := range groupNames {
		captureMap[currentGroupName] = matches[index]
	}
	return captureMap, true
}
