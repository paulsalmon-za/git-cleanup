package commands

import (
	"flag"
	"fmt"
	git "git-cleanup/git"
	dialog "git-cleanup/utils/dialog"
	hints "git-cleanup/utils/hints"
	"os"
	"strings"

	prompt "github.com/c-bata/go-prompt"
)

type protectCommand struct {
	list          bool
	help          bool
	defaultBranch bool
	flagSet       *flag.FlagSet
}

func (cmd *protectCommand) Initialize(flagSet *flag.FlagSet) {

	flagSet.BoolVar(&cmd.list, "list", false, "Display a list of protected branch paths")
	flagSet.BoolVar(&cmd.help, "help", false, "Show help for cleanup local")
	flagSet.BoolVar(&cmd.defaultBranch, "default", false, "Set as default branch")

	cmd.flagSet = flagSet
}

func (cmd *protectCommand) MinArgs() int {
	return 1
}

func (cmd *protectCommand) Execute(args []string) {
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

	if cmd.defaultBranch {

		if strings.Index("*", branchName) >= 0 {
			dialog.Error("No wildcard characters allowed for the default branch: %s", branchName)
			os.Exit(1)
		}

		localBranches := git.FindBranches(branchName)
		if len(localBranches) == 0 {
			dialog.Error("No local branch: %s", branchName)
			hints.CheckoutBranch(branchName)
			if dialog.PromptYesNo(fmt.Sprintf("Do you want to continue adding %s as the default branch?", branchName)) == false {
				os.Exit(1)
			}
		} else {
			localBranch := localBranches[0]
			if localBranch.Remote == "" {
				dialog.Error("No remote orgin for branch: %s", branchName)
				hints.CheckoutBranch(branchName)
				if dialog.PromptYesNo(fmt.Sprintf("Do you want to continue adding %s as the default branch?", branchName)) == false {
					os.Exit(1)
				}
			}
		}
		config.DefaultBranch = branchName
	}

	config.Protect(branchName)

	config.Save()

	dialog.Info("included %s in protected branches:", branchName)
	config.ShowProtectedPaths()
}

func (cmd *protectCommand) Help() {
	dialog.Help("cleanup protect:")
	cmd.flagSet.PrintDefaults()
}

func (cmd *protectCommand) Suggestions(in prompt.Document) []prompt.Suggest {
	branches := git.FindBranches("*")

	list := make([]prompt.Suggest, 0)
	for _, b := range branches {
		if b.Protected == false {
			description := ""
			if b.Active {
				description = "Active"
			}
			list = append(list, prompt.Suggest{Text: b.Name, Description: description})
		}
	}
	if strings.HasPrefix(in.GetWordBeforeCursor(), "-") == false {
		list = append(list, prompt.Suggest{Text: fmt.Sprintf("%s*", strings.TrimRight(in.GetWordBeforeCursor(), "*")), Description: "Add wildcard * to lock path"})
	}

	list = dialog.IncludeFlags(list, cmd.flagSet)

	return prompt.FilterHasPrefix(list, in.GetWordBeforeCursor(), true)
}
