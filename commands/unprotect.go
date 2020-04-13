package commands

import (
	"flag"
	"fmt"

	git "git-cleanup/git"
	dialog "git-cleanup/utils/dialog"

	prompt "github.com/c-bata/go-prompt"
)

type unprotectCommand struct {
	list bool
	help bool

	flagSet *flag.FlagSet
}

func (cmd *unprotectCommand) Initialize(flagSet *flag.FlagSet) {

	flagSet.BoolVar(&cmd.list, "list", false, "Display a list of protected branch paths")
	flagSet.BoolVar(&cmd.help, "help", false, fmt.Sprintf("Show help for cleanup %s", UNPROTECT))

	cmd.flagSet = flagSet
}

func (cmd *unprotectCommand) MinArgs() int {
	return 0
}

func (cmd *unprotectCommand) Execute(args []string) {

	if cmd.help {
		cmd.Help()
		return
	}

	config := git.GetConfig()

	if cmd.list {
		config.ShowProtectedPaths()

		return
	}

	branchName := args[2]

	config.RemoveProtected(branchName)
	if config.DefaultBranch == branchName {
		config.DefaultBranch = ""
	}
	config.Save()

	dialog.Info("removed %s from protected branches:", branchName)
	config.ShowProtectedPaths()
}

func (cmd *unprotectCommand) Help() {
	cmd.flagSet.PrintDefaults()
}

func (cmd *unprotectCommand) Suggestions(in prompt.Document) []prompt.Suggest {
	branches := git.FindBranches("*")

	list := make([]prompt.Suggest, 0)
	for _, b := range branches {
		if b.Protected == true {
			description := ""
			if b.Active {
				description = "Active"
			}
			list = append(list, prompt.Suggest{Text: b.Name, Description: description})
		}
	}

	list = dialog.IncludeFlags(list, cmd.flagSet)

	return prompt.FilterHasPrefix(list, in.GetWordBeforeCursor(), true)
}
