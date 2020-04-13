package commands

import (
	"flag"
	"fmt"
	"os"
	"strings"

	git "git-cleanup/git"
	dialog "git-cleanup/utils/dialog"
	hints "git-cleanup/utils/hints"

	prompt "github.com/c-bata/go-prompt"
)

type localCommand struct {
	force bool
	quiet bool
	help  bool

	flagSet *flag.FlagSet
}

func (cmd *localCommand) Initialize(flagSet *flag.FlagSet) {

	flagSet.BoolVar(&cmd.force, "force", false, "Do not prompt before applying cleanup")
	flagSet.BoolVar(&cmd.quiet, "quiet", false, "Prevent error when no branch is found")
	flagSet.BoolVar(&cmd.help, "help", false, "Show help for cleanup local")

	cmd.flagSet = flagSet
}

func (cmd *localCommand) MinArgs() int {
	return 0
}

func (cmd *localCommand) Execute(args []string) {
	if cmd.help {
		cmd.Help()
		return
	}

	hints.NoProtectedBranches()

	searchPath := "*"
	if len(args) > 2 {
		searchPath = args[2]

		fmt.Printf("Search for %s: ", searchPath)
	} else {
		fmt.Printf("Search for any branch: ")
	}

	if searchPath == "." {
		searchPath = "*"
	}

	branches := git.FindBranches(searchPath)
	available := branches.Available()
	notAvailable := branches.NotAvailable()

	if len(available) > 0 {
		fmt.Printf("Found the following branch(es):\n")
		available.Show()

		if cmd.force == false && dialog.PromptYesNo(fmt.Sprintf("\nAre you sure you want to remove (%d) local branch(es)?", len(available))) == false {
			os.Exit(0)
		}
		available.RemoveAll()

	} else {
		if len(notAvailable) > 0 {
			fmt.Printf("No available branches found\n\n")
			fmt.Printf("Found the following branches:\n")
			notAvailable.Show()

		} else {
			fmt.Printf("No branches found\n")
		}
		if cmd.quiet == false {
			os.Exit(1)
		}
	}
}

func (cmd *localCommand) Help() {
	cmd.flagSet.PrintDefaults()
}

func (cmd *localCommand) Suggestions(in prompt.Document) []prompt.Suggest {
	branches := git.FindBranches("*")

	list := make([]prompt.Suggest, 0)
	for _, b := range branches {
		description := b.Status()
		list = append(list, prompt.Suggest{Text: b.Name, Description: description})
	}

	list = append(list, prompt.Suggest{Text: ".", Description: "Any non protected branches"})

	if strings.HasPrefix(in.GetWordBeforeCursor(), "-") == false {
		wildcard := fmt.Sprintf("%s*", strings.TrimRight(in.GetWordBeforeCursor(), "*"))
		if wildcard != ".*" {
			list = append(list, prompt.Suggest{Text: wildcard, Description: "Wildcard search"})
		}
	}
	list = dialog.IncludeFlags(list, cmd.flagSet)

	return prompt.FilterHasPrefix(list, in.GetWordBeforeCursor(), true)
}
