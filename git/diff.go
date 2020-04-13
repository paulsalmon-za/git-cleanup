package git

import (
	"bufio"
	dialog "git-cleanup/utils/dialog"
	"strings"
)

// DiffFile ...
type DiffFile struct {
	Filename string
	Status   string
}

// Differences ...
type Differences []DiffFile

// Show ...
func (collecion Differences) Show() {
	for _, diff := range collecion {
		dialog.Info(diff.Filename)
	}
}

// Diff ...
func Diff() Differences {
	out := execGitOutPutOnlyIfErrorExists("diff", "--name-status", "--exit-code")

	scanner := bufio.NewScanner(strings.NewReader(out))

	list := make(Differences, 0)
	for scanner.Scan() {
		file := parseToDiffFile(scanner.Text())
		list = append(list, file)
	}

	return list
}

func parseToDiffFile(text string) DiffFile {
	splits := strings.Split(text, "\t")

	status := splits[0]

	name := strings.Join(splits[1:], " ")

	return DiffFile{Filename: name, Status: status}
}
