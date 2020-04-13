package main

import (
	"fmt"
	"os"
	"strings"

	commands "git-cleanup/commands"
	dialog "git-cleanup/utils/dialog"
	find "git-cleanup/utils/find"

	prompt "github.com/c-bata/go-prompt"
)

func commandSuggestions(in prompt.Document) []prompt.Suggest {
	textArgs := append([]string{"cleanup"}, strings.Split(in.Text, " ")...)

	if len(textArgs) > 2 {

		command := commands.GetCommand(textArgs...)

		if command != nil {
			return command.Suggestions(in)
		}
		return []prompt.Suggest{
			{Text: "invalid", Description: "There is no commands selected"},
		}
	}

	s := []prompt.Suggest{
		{Text: commands.LOCAL, Description: "Remove local branches"},
		// Next Version "remote pruning"
		{Text: commands.PROTECT, Description: "Prevents specified branches from being removed"},
		{Text: commands.UNPROTECT, Description: "Allow locked branches to be removed"},
		{Text: commands.RESET, Description: "Reset local repository to remote origin"},
	}

	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func showHelp() {

}

func askArgumentsIfNotProvided() []string {
	args := os.Args

	if len(os.Args) <= 1 {
		if find.ContainsString(args, "--help") {
			showHelp()
			os.Exit(0)
		}
		commandText := prompt.Input("git cleanup ", commandSuggestions)

		args = append([]string{"cleanup"}, strings.Split(commandText, " ")...)

	} else if len(os.Args) <= 2 {
		command := commands.GetCommand(os.Args...)
		commandName := os.Args[1]
		if command != nil {
			if find.ContainsString(args, "--help") {
				command.Help()
				os.Exit(0)
			}

			if len(os.Args)-2 < command.MinArgs() {
				commandText := prompt.Input(fmt.Sprintf("git cleanup %s ", commandName), command.Suggestions)

				args = append([]string{"cleanup", commandName}, strings.Split(commandText, " ")...)
			}

		}

	}

	return args
}

func main() {

	args := askArgumentsIfNotProvided()

	commandName := args[1]

	command := commands.GetCommand(args...)
	if command != nil {
		command.Execute(args)
	} else {
		dialog.Error("Could not determine command %s\n", commandName)
		showHelp()
		os.Exit(1)
	}
}
