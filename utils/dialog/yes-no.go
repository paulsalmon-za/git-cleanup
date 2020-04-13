package dialog

import (
	prompt "github.com/c-bata/go-prompt"
)

func yesNoCompleter(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "y", Description: "Yes"},
		{Text: "n", Description: "No"},
	}
	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

// PromptYesNo ...
func PromptYesNo(title string) bool {
	writer := prompt.NewStandardOutputWriter()
	displayQuestion(writer, title)
	displayNewline(writer)
	writer.Flush()

	in := prompt.Input("(y)es / (n)no: ", yesNoCompleter)

	if in == "yes" || in == "y" {
		return true
	}

	if in == "no" || in == "n" {
		return false
	}

	return PromptYesNo(title)
}
