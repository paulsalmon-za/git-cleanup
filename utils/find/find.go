package find

import (
	"regexp"
	"strings"
)

// RegexArray ...
type RegexArray []*regexp.Regexp

// GetMatchExpressions ...
func GetMatchExpressions(searchPaths ...string) RegexArray {
	result := make(RegexArray, 0)
	for _, s := range searchPaths {
		result = append(result, GetMatchExpression(s))
	}
	return result
}

// MatchAny ...
func (list RegexArray) MatchAny(value string) bool {
	if len(list) == 0 {
		return false
	}
	for _, r := range list {
		if r.MatchString(value) {
			return true
		}
	}

	return false

}

// GetMatchExpression ...
func GetMatchExpression(searchPath string) *regexp.Regexp {
	if searchPath == "" {
		searchPath = "*"
	}

	searchString := "(?i)^" + strings.Replace(searchPath, "*", ".*", -1) + "$"
	return regexp.MustCompile(searchString)
}

// ContainsString ...
func ContainsString(collection []string, value string) bool {
	if len(collection) == 0 {
		return false
	}
	for _, item := range collection {
		if item == value {
			return true
		}
	}

	return false
}
