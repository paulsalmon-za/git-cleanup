package git

import (
	"bufio"
	"strings"
)

// UntrackedFile ...
type UntrackedFile struct {
	Filename string
}

// UntrackedFiles ...
type UntrackedFiles []UntrackedFile

// Untracked ...
func Untracked() UntrackedFiles {
	out := execGit("ls-files", "-o", "--exclude-standard")

	scanner := bufio.NewScanner(strings.NewReader(out))

	list := make(UntrackedFiles, 0)
	for scanner.Scan() {
		file := parseToUntrackedFile(scanner.Text())
		list = append(list, file)
	}

	return list
}

func parseToUntrackedFile(text string) UntrackedFile {

	return UntrackedFile{Filename: text}
}
