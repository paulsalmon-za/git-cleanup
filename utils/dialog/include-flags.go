package dialog

import (
	"flag"

	prompt "github.com/c-bata/go-prompt"
)

// IncludeFlags ...
func IncludeFlags(collection []prompt.Suggest, flagSet *flag.FlagSet) []prompt.Suggest {
	result := make([]prompt.Suggest, 0)
	result = append(result, collection...)

	if flagSet == nil {
		return result
	}
	flagSet.VisitAll(func(f *flag.Flag) {
		suggestion := prompt.Suggest{Text: "-" + f.Name, Description: f.Usage}
		result = append(result, suggestion)
	})

	return result
}
