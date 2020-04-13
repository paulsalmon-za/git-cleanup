package git

import (
	"fmt"
	"os"
	"os/exec"
)

func execGit(params ...string) string {
	cmd := exec.Command("git", params...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("\n%s\n", out)
		os.Exit(1)
	}
	return string(out)
}

func execGitOutPutOnlyIfErrorExists(params ...string) string {
	cmd := exec.Command("git", params...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return string(out)
	}
	return ""
}

func execGitCheckError(params ...string) bool {
	cmd := exec.Command("git", params...)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return true
	}
	return false
}

// Fetch ...
func Fetch() string {
	out := execGit("fetch", "origin")

	return out
}

// Reset ...
func Reset() {
	execGit("reset", "--hard")
}

// Clean ...
func Clean() {
	execGit("clean", "-d", "-f")
}
