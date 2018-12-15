package retrieve

import (
	"regexp"
)

func extractGroupsAsMap(matches []string, regexp *regexp.Regexp) map[string]string {
	groups := make(map[string]string)
	matchNames := regexp.SubexpNames()
	for index, groupName := range matchNames {
		if index != 0 && groupName != "" {
			groups[groupName] = matches[index]
		}
	}
	return groups
}
