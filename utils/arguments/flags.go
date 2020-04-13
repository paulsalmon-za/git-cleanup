package arguments

import (
	"strings"
)

// OnlyFlags ...
func OnlyFlags(args []string) []string {
	list := make([]string, 0)
	foundFlag := false
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			foundFlag = true
		}

		if foundFlag {
			list = append(list, arg)
		}
	}

	return list
}
