package commands

import (
	"flag"
	"os"
	"strings"

	git "git-cleanup/git"
	dialog "git-cleanup/utils/dialog"
	hints "git-cleanup/utils/hints"

	prompt "github.com/c-bata/go-prompt"
)

type resetCommand struct {
	removeChanges   bool
	removeUntracked bool
	abortMerge      bool
	force           bool
	help            bool

	flagSet *flag.FlagSet
}

func (cmd *resetCommand) Initialize(flagSet *flag.FlagSet) {

	flagSet.BoolVar(&cmd.removeChanges, "remove-un-committed", false, "Remove any un-committed changes")
	flagSet.BoolVar(&cmd.removeUntracked, "remove-un-tracked", false, "Remove any un-tracked changes")
	flagSet.BoolVar(&cmd.abortMerge, "abort-merge", false, "Abort merge conflict in progress")
	flagSet.BoolVar(&cmd.force, "force", false, "Don't show any prompts")
	flagSet.BoolVar(&cmd.help, "help", false, "Show help for cleanup local")

	cmd.flagSet = flagSet
}

func (cmd *resetCommand) MinArgs() int {
	return 0
}

func (cmd *resetCommand) Execute(args []string) {
	if cmd.help {
		cmd.Help()
		return
	}

	hints.NoProtectedBranches()

	config := git.GetConfig()

	cmd.checkForMerge()
	cmd.checkForChanges()
	cmd.checkUnTrackedFiles()

	cmd.validateResetConfig(&config)

	git.Fetch()

	defaultBranchList := git.FindBranches(config.DefaultBranch)

	if len(defaultBranchList) == 0 { // Posibly try and checkout remote branch if it exists
		dialog.Error("No branch was found matching the default branch")
		dialog.Help("Make sure the default branch exists, run cleanup protect --list  to view protected / default branches")
		os.Exit(1)
	} else if len(defaultBranchList) > 1 {
		dialog.Error("The default branch matches multiple branches")
		dialog.Help("Do not use * wildcard characters when protecting branches with the --default flag")
		os.Exit(1)
	}
	defaultBranch := defaultBranchList[0]

	if defaultBranch.Remote == "" {
		dialog.Error("The default branch does not track a remote branch")
		hints.CheckoutBranch(defaultBranch.Name)
		os.Exit(1)
	}

	cmd.executeAbortMerge()
	cmd.executeRemoveUntracked()
	cmd.executeRemoveChanges()

	defaultBranch.Activate()

	branches := git.FindBranches("*")
	branches.Available().RemoveAll()

	defaultBranch.Reset()

	dialog.Info("Repository refreshed")
}

func (cmd *resetCommand) checkForMerge() {
	merge := git.GetMergeStatus()

	if merge.InProgress == false {
		return
	}

	dialog.Error("Merge conflict detected")

	if !cmd.abortMerge && !dialog.PromptYesNo("Do you want to abort the merge conflict and continue?") {
		os.Exit(1)
	}

	cmd.abortMerge = true
}

func (cmd *resetCommand) checkForChanges() {
	changes := git.Diff()

	if len(changes) == 0 {
		return
	}

	changes.Show()

	dialog.Error("There are %d un-committed changes", len(changes))

	if !cmd.removeChanges && !dialog.PromptYesNo("Do you want to remove the un-committed changes?") {
		os.Exit(1)
	}

	cmd.removeChanges = true
}

func (cmd *resetCommand) checkUnTrackedFiles() {
	untracked := git.Untracked()

	if len(untracked) == 0 {
		return
	}

	dialog.Info("There are %d un-tracked files", len(untracked))

	if cmd.removeUntracked == false {
		cmd.removeUntracked = dialog.PromptYesNo("Do you want to remove the un-tracked files?")
	}
}

func (cmd *resetCommand) validateResetConfig(cfg *git.Config) {
	if strings.Trim(cfg.DefaultBranch, " ") == "" {
		dialog.Error("No default branch specified")
		command := getCommandByName(PROTECT)
		command.Help()
		os.Exit(1)
	}

	if len(cfg.Protected) == 0 {
		dialog.Error("No protected branches specified")
		command := getCommandByName(PROTECT)
		command.Help()
		os.Exit(1)
	}
}

func (cmd *resetCommand) executeAbortMerge() {
	if cmd.abortMerge == false {
		return
	}
	git.AbortMerge()
}

func (cmd *resetCommand) executeRemoveChanges() {
	if cmd.removeChanges == false {
		return
	}
	git.Reset()
}

func (cmd *resetCommand) executeRemoveUntracked() {
	if cmd.removeUntracked == false {
		return
	}
	git.Clean()
}

func (cmd resetCommand) Help() {
	cmd.flagSet.PrintDefaults()
}

func (cmd *resetCommand) Suggestions(in prompt.Document) []prompt.Suggest {

	list := make([]prompt.Suggest, 0)

	list = dialog.IncludeFlags(list, cmd.flagSet)

	return prompt.FilterHasPrefix(list, in.GetWordBeforeCursor(), true)
}
