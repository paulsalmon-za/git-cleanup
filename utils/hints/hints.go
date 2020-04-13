package hints

import (
	git "git-cleanup/git"
	dialog "git-cleanup/utils/dialog"
)

// CheckoutBranch ...
func CheckoutBranch(branchName string) {
	dialog.Help("git fetch")
	dialog.Help("git checkout %s", branchName)
	dialog.Help("https://stackoverflow.com/questions/1783405/how-do-i-check-out-a-remote-git-branch")
}

// NoProtectedBranches ...
func NoProtectedBranches() {
	config := git.GetConfig()

	if len(config.Protected) > 0 {
		return
	}

	dialog.Error("You don't have any protected branches specified")
	dialog.Help("Protecting branches prevents branches from being removed")
	dialog.Help("git-cleanup protect pattern")
	dialog.Help("Patterns can contain * wildcard characters")
}
